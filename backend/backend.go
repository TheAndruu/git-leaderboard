package backend

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"os"
	"path/filepath"

	"github.com/TheAndruu/git-leaderboard/models"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var templates = make(map[string]*template.Template)

func init() {
	initializeTemplates()
	defineRoutes()
}

func defineRoutes() {
	http.HandleFunc("/repostats", saveRepoPost)
	http.HandleFunc("/recently-submitted", showRecentlySubmitted)
	http.HandleFunc("/most-commits", showMostCommits)
	http.HandleFunc("/", showHome)
}

// Base template is 'theme.html'  Can add any variety of content fillers in /layouts directory
func initializeTemplates() {
	layouts, err := filepath.Glob("static/templates/*.html")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Issue setting up templates ", err)
	}

	for _, layout := range layouts {
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(layout, "static/templates/layouts/theme.html"))
	}
}

func saveRepoPost(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "Saving repo stats")
	target := models.RepoStats{}
	json.NewDecoder(r.Body).Decode(&target)
	defer r.Body.Close()

	// TODO: Add validation on the fields of the struct - strip out anything not valid

	// TODO:
	// add fields for the numTotalCommits, numLeadAuthor, percentLeadAuthorToTotal
	computeStats(ctx, &target)

	// Save the data
	log.Infof(ctx, fmt.Sprintf("Accepting repo name: %v", target.RepoName))
	id, err := SaveStats(ctx, &target)
	if err != nil {
		log.Errorf(ctx, "Issue saving stats to datastore %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(420)
		values := map[string]string{"message": "Issue saving git stats to database"}
		asBytes, _ := json.Marshal(values)
		w.Write(asBytes)
		return
	}

	// Write content-type, statuscode, payload
	log.Infof(ctx, fmt.Sprintf("Created stats with id %v", id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	values := map[string]string{"message": fmt.Sprintf("Successfully stored stats for %s", target.RepoName)}
	asBytes, _ := json.Marshal(values)
	w.Write(asBytes)
}

func computeStats(ctx context.Context, stats *models.RepoStats) {

	var totalCommits int
	var authorCount int
	var leadAuthorTotal int

	for _, stat := range stats.Commits {
		// Sum the total number of commits
		totalCommits += stat.NumCommits
		// Sum the total number of authors
		authorCount++
		// Ensure we set the max commit value
		if stat.NumCommits > leadAuthorTotal {
			leadAuthorTotal = stat.NumCommits
		}
	}

	// calculate mean and percent
	totalCommitFloat := float64(totalCommits)
	leadAuthorPercent := float64(leadAuthorTotal) / totalCommitFloat
	averageAuthorCommits := totalCommitFloat / float64(authorCount)

	// calculate standard deviation
	// https://www.mathsisfun.com/data/standard-deviation-formulas.html
	var sumOfSquaredDifferencesFromMean float64
	for _, stat := range stats.Commits {
		diffFromMean := float64(stat.NumCommits) - averageAuthorCommits
		sumOfSquaredDifferencesFromMean += math.Pow(diffFromMean, 2)
	}
	meanOfSquaredDifferences := sumOfSquaredDifferencesFromMean / float64(authorCount)
	commitDeviation := math.Sqrt(meanOfSquaredDifferences)

	// set the values on the struct
	stats.AuthorCount = authorCount
	stats.AverageAuthorCommits = averageAuthorCommits
	stats.LeadAuthorPercent = leadAuthorPercent
	stats.LeadAuthorTotal = leadAuthorTotal
	stats.TotalCommits = totalCommits
	stats.CommitDeviation = commitDeviation
	stats.CoefficientVariation = commitDeviation / averageAuthorCommits

	if len(stats.Commits) > 10 {
		stats.Commits = stats.Commits[:10]
	}

	// TODO: add coefficient of variation:
	// https://en.wikipedia.org/wiki/Coefficient_of_variation
}
