package backend

import (
	"net/http"

	"google.golang.org/appengine"
)

// A Welcome message with title, demonstrates passing data to a template
type Welcome struct {
	Title   string
	Message string
}

// A template taking a struct pointer (&message) containing data to render
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	message := Welcome{Title: "Bootstrap, Go, and GAE", Message: "Bootstrap added to Golang on App Engine.  Feel free to customize further"}

	// outerTheme refernces the template defined within theme.html
	templates["welcome.html"].ExecuteTemplate(w, "outerTheme", &message)
}

// Shows the leaders in the current git repos
func showLeaders(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// get the 10 most recent RepoStats
	recentStats := GetRecentRepoStats(ctx, 10)

	// outerTheme refernces the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", &recentStats)
}
