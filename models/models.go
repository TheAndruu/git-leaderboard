package models

// RepoStats contains stats for a git repo
type RepoStats struct {
	RepoName string
	RepoURL  string
	Commits  []CommitCount
}

// CommitCount represents a count and author
type CommitCount struct {
	Author     string
	NumCommits int
}
