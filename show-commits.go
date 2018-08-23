package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	execShortlog()
}

func execShortlog() {
	cmdOut, err := exec.Command("git", "shortlog", "master", "--summary", "--numbered").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running the git command: ", err)
		os.Exit(1)
	}

	commitLines := strings.Split(string(cmdOut), "\n")

	for _, element := range commitLines {
		if len(element) < 1 {
			// Any line without a report in it (separator, last line)
			continue
		}
		commitLine := strings.Split(element, "\t")
		fmt.Fprintln(os.Stdout, "Number of elements in commitLine: ", len(commitLine))
		fmt.Printf("%s has %s commits\n", commitLine[1], commitLine[0])
	}

	// fmt.Printf("Output: \n%s\n", output)
}
