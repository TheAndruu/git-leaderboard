package backend

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/TheAndruu/git-leaderboard/models"
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
	http.HandleFunc("/", showLeaders)
}

// Base template is 'theme.html'  Can add any variety of content fillers in /layouts directory
func initializeTemplates() {
	layouts, err := filepath.Glob("../frontend/templates/*.html")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Issue setting up templates ", err)
	}

	for _, layout := range layouts {
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(layout, "../frontend/templates/layouts/theme.html"))
	}
}

func saveRepoPost(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "Saving repo stats")
	target := models.RepoStats{}
	json.NewDecoder(r.Body).Decode(&target)
	defer r.Body.Close()

	// TODO: Add validation on the fields of the struct - strip out anything not valid

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
