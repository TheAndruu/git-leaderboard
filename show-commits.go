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
	repoStats := getRepoOriginsFromGit()
	repoStats.Commits = getRepoCommits()
	submitRepoStats(&repoStats)

	// TODO: Show URL of the leaderboard UI
}

/** Queries git to determine the name and remote url of the repo */
func getRepoOriginsFromGit() models.RepoStats {

	// could also use basename here, but fear it wouldn't be on everyone's machine
	repoNameOut, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Please run this command from within an existing git repository")
		fmt.Fprintln(os.Stderr, "And ensure you have git installed")
		fmt.Fprintln(os.Stderr, "There was an error running the git command: ", err)
		os.Exit(1)
	}

	// lob off everything up to and including the last slash
	repoNameStr := strings.TrimSpace(string(repoNameOut))
	var splitName []string
	if strings.Contains(repoNameStr, "/") {
		splitName = strings.Split(repoNameStr, "/")
	} else {
		splitName = strings.Split(repoNameStr, "\\")
	}
	repoNameStr = splitName[len(splitName)-1]

	//Now get and supply repoUrl
	repoURLOut, err := exec.Command("git", "remote", "get-url", "--push", "origin").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Please ensure the git repo has an upstream origin")
		fmt.Fprintln(os.Stderr, "There was an issue getting the remote git URL: ", err)
		os.Exit(2)
	}
	repoURLStr := strings.TrimSpace(string(repoURLOut))
	stats := models.RepoStats{RepoName: repoNameStr, RepoURL: repoURLStr}
	return stats
}

func getRepoCommits() []models.CommitCount {
	cmdOut, err := exec.Command("git", "shortlog", "master", "--summary", "--numbered").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error checking the commit stats: ", err)
		os.Exit(3)
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
			fmt.Fprintln(os.Stderr, "Trouble parsing the number of commits for author: ", author, err)
			os.Exit(4)
		}
		authorCommit := models.CommitCount{Author: author, NumCommits: numCommits}
		commitCounts = append(commitCounts, authorCommit)
	}
	return commitCounts
}

func submitRepoStats(repoStats *models.RepoStats) {
	fmt.Println("Submitting stats to leaderboard")
	fmt.Println("Project name: " + repoStats.RepoName)
	fmt.Println("Project URL: " + repoStats.RepoURL)
	url := "https://backend-gl.appspot.com/repostats"
	//url := " http://localhost:8080/repostats"

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
