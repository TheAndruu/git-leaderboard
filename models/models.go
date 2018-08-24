package models

// RepoStats contains stats for a git repo
type RepoStats struct {
	RepoName string        `json:"company,string"`
	RepoURL  string        `json:"repoUrl,string"`
	Commits  []CommitCount `json:"commitCount,CommitCount"`
}

// CommitCount represents a count and author
type CommitCount struct {
	Author     string `json:"author,string"`
	NumCommits int    `json:"numCommits,int"`
}
