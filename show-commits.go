package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/TheAndruu/git-leaderboard/models"
)

func main() {
	// TODO: Exit and show error message if run from non-git directory
	repoStats := getRepoOriginsFromGit()
	repoStats.Commits = getRepoCommits()
	submitRepoStats(&repoStats)
}

/** Queries git to determine the name and remote url of the repo */
func getRepoOriginsFromGit() models.RepoStats {

	repoURL := "www.foo"

	// could also use basename here, but fear it wouldn't be on everyone's machine
	repoNameOut, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Please run this command from within an existing git repository")
		fmt.Fprintln(os.Stderr, "And ensure you have git installed")
		fmt.Fprintln(os.Stderr, "There was an error running the git command: ", err)
		os.Exit(2)
	}

	// lob off everything up to and including the last slash
	repoNameStr := string(repoNameOut)
	var splitName []string
	if strings.Contains(repoNameStr, "/") {
		splitName = strings.Split(repoNameStr, "/")
	} else {
		splitName = strings.Split(repoNameStr, "\\")
	}
	repoNameStr = splitName[len(splitName)-1]

	// fmt.Println(fmt.Sprintf("Got: %s", repoNameStr))
	// os.Exit(9)

	// TODO: now get and supply repoUrl

	stats := models.RepoStats{RepoName: repoNameStr, RepoURL: repoURL}
	return stats
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

	shortLogString := string(cmdOut)
	fmt.Println(fmt.Sprintf("Stats from the git repo, nice job! \n%s", shortLogString))

	commitLines := strings.Split(shortLogString, "\n")

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

func submitRepoStats(repoStats *models.RepoStats) {
	fmt.Println("Got here again")

	url := "https://backend-gl.appspot.com/repostats"
	//url := " http://localhost:8080/repostats"

	fmt.Println("URL:>", url)

	jsonValue, _ := json.Marshal(repoStats)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
