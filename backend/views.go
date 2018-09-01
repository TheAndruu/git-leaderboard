package backend

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// Shows the leaders in the current git repos
func showLeaders(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// get the 10 most recent RepoStats
	recentStats := GetRecentRepoStats(ctx, 10)

	log.Infof(ctx, "Got the stats %v", len(*recentStats))

	// outerTheme refernces the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", &recentStats)
}
