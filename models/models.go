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
	TotalCommits int `json:"totalCommits"`
	// Number of authors
	AuthorCount     int `json:"authorCount"`
	LeadAuthorTotal int `json:"leadAuthorTotal"`
	// Percent out of total number of commits
	LeadAuthorPercent float64 `json:"leadAuthorPercent"`
	// Average number of commits per author
	AverageAuthorCommits float64 `json:"averageAuthorCommits"`
	// The standard deviation of commits
	// https://www.mathsisfun.com/data/standard-deviation.html
	CommitDeviation float64 `json:"commitDeviation"`
	// https://en.wikipedia.org/wiki/Coefficient_of_variation
	CoefficientVariation float64 `json:"coefficientVariation"`
}

// CommitCount represents a count and author
type CommitCount struct {
	Author     string `json:"author"`
	NumCommits int    `json:"numCommits"`
}
