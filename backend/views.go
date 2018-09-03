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

	// outerTheme references the template defined within theme.html
	templates["home.html"].ExecuteTemplate(w, "outerTheme", pageData)
}

// Shows the leaders in the current git repos
func showRecentlySubmitted(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// get the 10 most recent RepoStats
	recentStats := GetStatsOrderedBy(ctx, 2, "-DateUpdated", fetchLimit)

	pageData := RepoStatsPage{
		Title:          "Recently Submitted",
		SubTitle:       "Top committers in the most recently-submitted repositories",
		ContentHead:    "Latest Projects",
		ContentMessage: "These are the projects that have been submitted most recently. Only considers projects with at least 2 authors. ",
		MenuSection:    "recently-submitted",
		RepoStats:      recentStats}

	// outerTheme references the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", pageData)
}

// Shows the leaders based on most overall commits
func showMostCommits(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// get the 10 most recent RepoStats
	recentStats := GetStatsOrderedBy(ctx, 2, "-TotalCommits", fetchLimit)

	pageData := RepoStatsPage{
		Title:          "Most Commits",
		SubTitle:       "Highest Overall Commits per Project",
		ContentHead:    "Highest number of overall commits",
		ContentMessage: "These projects are those which have the highest number of commits by all authors added together. Only considers projects with at least 2 authors. ",
		MenuSection:    "most-commits",
		RepoStats:      recentStats}

	// outerTheme references the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", pageData)
}

// Shows the leaders based on most authors
func showMostAuthors(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// get the 10 most recent RepoStats
	recentStats := GetStatsOrderedBy(ctx, 2, "-AuthorCount", fetchLimit)

	pageData := RepoStatsPage{
		Title:          "Most Authors",
		SubTitle:       "Projects with the Most Authors",
		ContentHead:    "More code committers than any other project",
		ContentMessage: "Projects which have the highest number of authors submitting code updates. Only considers projects with at least 2 authors. ",
		MenuSection:    "most-authors",
		RepoStats:      recentStats}

	// outerTheme references the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", pageData)
}

// Shows the leaders based on the most commits by a single author
func showMostSingleAuthor(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// get the 10 most recent RepoStats
	recentStats := GetStatsOrderedBy(ctx, 2, "-LeadAuthorTotal", fetchLimit)

	pageData := RepoStatsPage{
		Title:          "Biggest Single Author",
		SubTitle:       "Most Commits by a Single Author",
		ContentHead:    "Projects with biggest lead author",
		ContentMessage: "The authors of these repositories have the most commits out of any other. Only considers projects with at least 2 authors. ",
		MenuSection:    "most-single-author",
		RepoStats:      recentStats}

	// outerTheme references the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", pageData)
}

// Shows the leaders based on author with the highest percent
func showLeadAuthorHighestPercent(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// get the 10 most recent RepoStats
	recentStats := GetStatsOrderedBy(ctx, 10, "-LeadAuthorPercent", fetchLimit)

	pageData := RepoStatsPage{
		Title:          "Authors with Highest Percentage",
		SubTitle:       "Lead Authors with Highest Percentage of Commits",
		ContentHead:    "Leading Commit Percentage",
		ContentMessage: "Projects whose leading committer have the highest percentage of commits compared to the overall number of commits. Only considers projects with at least 10 authors. ",
		MenuSection:    "lead-author-highest-percent",
		RepoStats:      recentStats}

	// outerTheme references the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", pageData)
}

// Shows the leaders based on authors having the highest average commits
func showHighestAverageCommits(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// get the 10 most recent RepoStats
	recentStats := GetStatsOrderedBy(ctx, 10, "-AverageAuthorCommits", fetchLimit)

	pageData := RepoStatsPage{
		Title:          "Average Commits",
		SubTitle:       "Highest Average Commits by Author",
		ContentHead:    "Projects with Highest Commit Average",
		ContentMessage: "The authors on these projects enjoy the highest number of average commits among all projects. Only considers projects with at least 10 authors. ",
		MenuSection:    "highest-average-commits",
		RepoStats:      recentStats}

	// outerTheme references the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", pageData)
}

// Shows the projects with least standard deviation between authors
func showLowestStandardDeviation(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// get the 10 most recent RepoStats
	recentStats := GetStatsOrderedBy(ctx, 10, "CommitDeviation", fetchLimit)

	pageData := RepoStatsPage{
		Title:          "Least Standard Deviation",
		SubTitle:       "Lowest Standard Deviation of Commits by Author",
		ContentHead:    "Authors With Least Standard Deviation in Commits",
		ContentMessage: "These projects' authors feature the least standard deviation among their commit counts. Only considers projects with at least 10 authors. ",
		MenuSection:    "lowest-standard-deviation",
		RepoStats:      recentStats}

	// outerTheme references the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", pageData)
}

// Shows the projects with least coefficient variation
func showLeastCoefficientVariation(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// get the 10 most recent RepoStats
	recentStats := GetStatsOrderedBy(ctx, 10, "CoefficientVariation", fetchLimit)

	pageData := RepoStatsPage{
		Title:          "Least Coefficient Variation",
		SubTitle:       "Lowest Variance Coefficient among Authors",
		ContentHead:    "Least Variance in Commits by Authors",
		ContentMessage: "Authors in these projects have the least variance in the number of commits they contribute.  This is the best metric to see prjoects where multiple authors share the load. Only considers projects with at least 10 authors. ",
		MenuSection:    "least-coefficient-variation",
		RepoStats:      recentStats}

	// outerTheme references the template defined within theme.html
	templates["leaderboard.html"].ExecuteTemplate(w, "outerTheme", pageData)
}
