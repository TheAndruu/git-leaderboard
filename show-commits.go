package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/kr/pretty"

	"github.com/TheAndruu/git-leaderboard/models"
)

func main() {
	repoStats := initStatsFromArgs()
	repoStats.Commits = getRepoCommits()
	pretty.Print(repoStats)
}

func initStatsFromArgs() models.RepoStats {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Please enter 2 parameters: the project name and its URL, for example:")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "    git-leaderboard \"golongpoll\" \"https://github.com/jcuga/golongpoll\"")
		fmt.Fprintln(os.Stderr, "")
		os.Exit(1)
	}
	stats := models.RepoStats{RepoName: args[0], RepoURL: args[1]}
	return stats
}

func getRepoCommits() []models.CommitCount {
	cmdOut, err := exec.Command("git", "shortlog", "master", "--summary", "--numbered").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running the git command: ", err)
		os.Exit(2)
	}

	commitLines := strings.Split(string(cmdOut), "\n")

	var commitCounts []models.CommitCount

	for _, element := range commitLines {
		if len(element) < 1 {
			// Any line without a report in it (separator, last line)
			continue
		}
		commitLine := strings.Split(element, "\t")
		author := strings.TrimSpace(commitLine[1])
		numCommits, err := strconv.Atoi(strings.TrimSpace(commitLine[0]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Trouble reading the number of commits for author: ", author, err)
			os.Exit(3)
		}
		authorCommit := models.CommitCount{Author: author, NumCommits: numCommits}
		commitCounts = append(commitCounts, authorCommit)
	}
	return commitCounts
}
