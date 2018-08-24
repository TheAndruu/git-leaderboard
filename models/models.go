package models

// RepoStats contains stats for a git repo
type RepoStats struct {
	RepoName string        `json:"repoName"`
	RepoURL  string        `json:"repoUrl"`
	Commits  []CommitCount `json:"commits"`
}

// CommitCount represents a count and author
type CommitCount struct {
	Author     string `json:"author"`
	NumCommits int    `json:"numCommits"`
}
