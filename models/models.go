package models

import "time"

// RepoStats contains stats for a git repo
type RepoStats struct {
	ID         string        `json:"id"`
	RepoName   string        `json:"repoName"`
	RepoURL    string        `json:"repoUrl"`
	Commits    []CommitCount `json:"commits"`
	DateUpated time.Time     `json:"dateUpdated"`
}

// CommitCount represents a count and author
type CommitCount struct {
	Author     string `json:"author"`
	NumCommits int    `json:"numCommits"`
}
