# git-leaderboard

Share projects' git stats and compete on the leaderboard!

View the centralized leaderboard at: [Git Commit Leaderboard](https://backend-gl.appspot.com/)

## What the client executable does

The client executable lists all authors with commits in the current repository and the number of commits each has, ranked from high to low.  

It then publishes the name of the git repo, it's remote origin, and the author commit stats to the [Git Commit Leaderboard](https://backend-gl.appspot.com/).

## Pre-requisites

Must have 'git' installed on the machine

**Note on the instructions below:**

Rather than copy the executables to the git repository, consider moving it to a folder on your $PATH, such as `sudo mv ~/Downloads/show-commits /usr/bin`.  This way you can run `show-commits` from anywhere on your machine.

## Running on Linux

Download the [Linux executable](https://github.com/TheAndruu/git-leaderboard/raw/master/build/show-commits)

Then run it within a git repository on your local machine.  You may need to make the downloaded binary file executable with: `chmod +x show-commits`

For example, if the file was downloaded to your `~/Downloads` folder, run:

    cd ~/git/some-git-project-to-analyze
    cp ~/Downloads/show-commits .
    chmod +x show-commits
    ./show-commits

## Running on Mac OSX

Download the [Mac OSX executable](https://github.com/TheAndruu/git-leaderboard/raw/master/build/show-commits-mac)

Then run it within a git repository on your local machine.  You may need to make the downloaded binary file executable with: `chmod +x show-commits`

For example, if the file was downloaded to your `~/Downloads` folder, run:

    cd ~/git/some-git-project-to-analyze
    cp ~/Downloads/show-commits-mac .
    chmod +x show-commits-mac
    ./show-commits-mac

## Running on Windows

Download the [Windows executable](https://github.com/TheAndruu/git-leaderboard/raw/master/build/show-commits.exe)

Then copy `show-commits.exe` to a git repository on your local machine.  

Open a command prompt to that git repository and execute: `show-commits.exe`.

## Running from go

If you have golang installed on your local machine, you can execute:

     go get github.com/TheAndruu/git-leaderboard
     cd $GOPATH/src/github.com/TheAndruu/git-leaderboard
     go install show-commits.go

If you have $GOBIN set up on your $PATH, you can now run `show-commits` from any git repo on your machine.

### Notes

To compile across systems, needed to first execute these commands to install all the windows-amd64 standard packages on your system:

     GOOS=linux GOARCH=amd64 go install
     GOOS=darwin GOARCH=amd64 go install
     GOOS=windows GOARCH=amd64 go install
