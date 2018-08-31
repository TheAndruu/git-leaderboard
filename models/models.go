package models

import "time"

// RepoStats contains stats for a git repo
type RepoStats struct {
	RepoName    string        `json:"repoName"`
	RepoURL     string        `json:"repoUrl"`
	Commits     []CommitCount `json:"commits"`
	DateUpdated time.Time     `json:"dateUpdated"`

	//
	// Server-side stats:
	TotalCommits    int `json:"totalCommits"`
	LeadAuthorTotal int `json:"leadAuthorRotal"`
	// Percent out of total number of commits
	LeadAuthorPercent int `json:"leadAuthorPercent"`
	// Average number of commits per author
	AverageAuthorCommits int `json:"averageAuthorCommits"`
	// The standard deviation of commits
	// https://www.mathsisfun.com/data/standard-deviation.html
	CommitDeviation int `json:"commitDeviation"`
}

// CommitCount represents a count and author
type CommitCount struct {
	Author     string `json:"author"`
	NumCommits int    `json:"numCommits"`
}
