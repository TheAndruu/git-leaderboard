# git-leaderboard

Share projects' git stats and compete on the leaderboard!

View the centralized leaderboard at: [Git Commit Leaderboard](http://backend-gl.appspot.com/)

## To submit stats for a repo

It's easy as downloading the binary and running it from any local git repo on your machine.

### Pre-requisites

Must have 'git' installed on the machine

### To run

Download the binary executable [show-commits](https://github.com/TheAndruu/git-leaderboard/raw/master/show-commits)

Open a Terminal session to the root of a git repository and execute the downloaded file.

Repo stats will be printed on the screen and submitted to the central leaderboard.

Alternatively, if you have golang installed on your local machine, you can execute:

     go get github.com/TheAndruu/git-leaderboard

Then navigate to the directory where you fetched the code and install the project  `go install show-commits.go`.  If you have $GOBIN set up correctly, you should be able to run `show-commits` from any git repo on your machine.
