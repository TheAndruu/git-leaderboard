# git-leaderboard

Share projects' git stats and compete on the leaderboard!

View the centralized leaderboard at: [Git Commit Leaderboard](https://backend-gl.appspot.com/)

## To submit stats for a repo

It's easy as downloading the binary and running it from any local git repo on your machine.

## Pre-requisites

Must have 'git' installed on the machine

## Running on linux

* [Linux](https://github.com/TheAndruu/git-leaderboard/raw/master/build/show-commits)

Then execute the file from a git repository on your local machine.  You may need to make the downloaded binary file executable with: `chmod +x show-commits`

For example, if the file was downloaded to your `~/Downloads` folder on Linux:

    cd ~/git/some-git-project-to-analyze
    cp ~/Downloads/show-commits .
    chmod +x show-commits
    ./show-commits

## Running on Mac OSX

* [MacOSX](https://github.com/TheAndruu/git-leaderboard/raw/master/build/show-commits-mac)

## Running on Windows

* [Windows](https://github.com/TheAndruu/git-leaderboard/raw/master/build/show-commits.exe)


Repo stats will be printed on the screen and submitted to the central leaderboard.

Instead of copying `show-commits` to your local git repo, consider dropping it in a folder on your $PATH, such as `sudo mv ~/Downloads/show-commits /usr/bin`.  This way you can run `show-commits` from anywhere on your machine.

### To run from go

Alternatively, if you have golang installed on your local machine, you can execute:

     go get github.com/TheAndruu/git-leaderboard
     cd $GOPATH/src/github.com/TheAndruu/git-leaderboard
     go install show-commits.go

If you have $GOBIN set up on your $PATH, you can now run `show-commits` from any git repo on your machine.
