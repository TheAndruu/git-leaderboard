package backend

import (
	"net/http"

	"github.com/TheAndruu/git-leaderboard/models"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// RepoStatsPage has a title, subtitle, and list of RepoStats for display
type RepoStatsPage struct {
	Title       string
	SubTitle    string
	MenuSection string
	RepoStats   *[]models.RepoStats
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
	recentStats := GetRecentRepoStats(ctx, 10)
	log.Infof(ctx, "Got the stats %v", len(*recentStats))

	pageData := RepoStatsPage{
		Title:       "Latest submissions",
		SubTitle:    "Top committers in the most recently-submitted repositories",
		MenuSection: "recently-submitted",
		RepoStats:   recentStats}

	// outerTheme refernces the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", pageData)
}
