package backend

import (
	"net/http"

	"github.com/TheAndruu/git-leaderboard/models"
	"google.golang.org/appengine"
)

var fetchLimit = 50

// RepoStatsPage has a title, subtitle, and list of RepoStats for display
type RepoStatsPage struct {
	Title          string
	SubTitle       string
	MenuSection    string
	ContentHead    string
	ContentMessage string
	RepoStats      *[]models.RepoStats
}

// Shows the leaders in the current git repos
func showHome(w http.ResponseWriter, r *http.Request) {
	pageData := RepoStatsPage{
		Title:       "Git Leaderboard",
		SubTitle:    "Share projects' git stats and compete on the leaderboard!",
		MenuSection: "home"}

	// outerTheme refernces the template defined within theme.html
	templates["home.html"].ExecuteTemplate(w, "outerTheme", pageData)
}

// Shows the leaders in the current git repos
func showRecentlySubmitted(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// get the 10 most recent RepoStats
	recentStats := GetStatsOrderedBy(ctx, "-DateUpdated", fetchLimit)

	pageData := RepoStatsPage{
		Title:          "Recently Submitted",
		SubTitle:       "Top committers in the most recently-submitted repositories",
		ContentHead:    "Latest Projects",
		ContentMessage: "These are the projects that have been submitted most recently.",
		MenuSection:    "recently-submitted",
		RepoStats:      recentStats}

	// outerTheme refernces the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", pageData)
}

// Shows the leaders based on most overall commits
func showMostCommits(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// get the 10 most recent RepoStats
	recentStats := GetStatsOrderedBy(ctx, "-TotalCommits", fetchLimit)

	pageData := RepoStatsPage{
		Title:          "Most Commits",
		SubTitle:       "Projects with the most commits overal",
		ContentHead:    "Highest number of overall commits",
		ContentMessage: "These projects are those which have the highest number of commits by all authors added together.",
		MenuSection:    "most-commits",
		RepoStats:      recentStats}

	// outerTheme refernces the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", pageData)
}
