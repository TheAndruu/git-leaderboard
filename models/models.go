package models

// RepoStats contains stats for a git repo
type RepoStats struct {
	RepoName string
	RepoURL  string
	Commits  []CommitTally
}

// CommitTally represents a count and author
type CommitTally struct {
	Author     string
	NumCommits int
}
