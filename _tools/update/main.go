package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"bitbucket.org/creachadair/shell"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"gopkg.in/src-d/go-errors.v1"
	"gopkg.in/src-d/go-log.v1"
)

const (
	org          = "bblfsh"
	repo         = "documentation"
	repoURL      = "https://github.com/" + org + "/" + repo
	commitMsg    = "regular languages update"
	branchPrefix = "auto-update-languages"

	gitUser = "bblfsh-release-bot"
	gitMail = "<release-bot@bblf.sh>"

	errSpecialText = "nothing to commit"
)

var (
	errCmdFailed = errors.NewKind("command failed: %v, output: %v")

	errFailedToPrepareBranch = errors.NewKind("failed to prepare branch")
	errFailedToPreparePR     = errors.NewKind("failed to prepare pull request")
	errNothingToCommit       = errors.NewKind(errSpecialText)
)

type pipeLine struct {
	nodes []pipeLineNode
}

type pipeLineNode struct {
	logFormat string
	command   string
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	branch := getBranch()

	pipeLine := newPipeLine(token, branch, commitMsg)
	if err := pipeLine.exec(); err != nil {
		if errNothingToCommit.Is(err) {
			log.Infof("no changes detected")
			return
		}
	}

	if err := preparePR(token, branch, commitMsg); err != nil {
		log.Infof(err.Error())
		os.Exit(1)
	}
}

// preparePR creates pull request
func preparePR(githubToken, branch, commitMsg string) error {
	ctx := context.Background()
	client := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)))

	log.Infof("Preparing pr %v -> master", branch)
	newPR := &github.NewPullRequest{
		Title:               &branch,
		Head:                &branch,
		Base:                strPtr("master"),
		Body:                strPtr(commitMsg),
		MaintainerCanModify: newTrue(),
	}

	pr, _, err := client.PullRequests.Create(ctx, org, repo, newPR)
	if err != nil {
		return errFailedToPreparePR.Wrap(err)
	}

	log.Infof("pull request %v has been successfully created", *pr.ID)
	return nil
}

func newPipeLine(githubToken, branch, commitMsg string) (result pipeLine) {
	var nodes []pipeLineNode
	nodes = append(nodes,
		pipeLineNode{
			logFormat: "performing git clone repository",
			command:   fmt.Sprintf("git remote rm origin ; git remote add origin %s", getOrigin(githubToken)),
		}, pipeLineNode{
			logFormat: "creating branch %v",
			command:   fmt.Sprintf("git checkout -b %s", shell.Quote(branch)),
		},
		pipeLineNode{
			logFormat: "set git user info",
			command:   fmt.Sprintf("git config --global user.name %v ; git config --global user.email %v", gitUser, shell.Quote(gitMail)),
		}, pipeLineNode{
			logFormat: "committing the changes",
			command:   fmt.Sprintf("git add -A ; git commit --signoff -m \"%s\"", commitMsg),
		}, pipeLineNode{
			logFormat: "pushing changes",
			command:   fmt.Sprintf("git push origin %s", branch),
		},
	)

	return pipeLine{nodes}
}

func (p pipeLine) exec() error {
	for _, c := range p.nodes {
		log.Infof(c.logFormat)
		if err := execCmd(c.command); err != nil {
			err = errFailedToPrepareBranch.Wrap(err)
			if strings.Contains(err.Error(), errSpecialText) {
				err = errNothingToCommit.Wrap(err)
			}
		}
	}
	return nil
}

// execCmd executes the specified Bash script. If execution fails, the error contains
// the combined output from stdout and stderr of the script.
// Do not use this for scripts that produce a large volume of output.
func execCmd(command string) error {
	cmd := exec.Command("bash", "-c", command)

	data, err := cmd.CombinedOutput()
	log.Debugf("command output: %v", string(data))
	if err != nil {
		return errCmdFailed.New(err, string(data))
	}

	return nil
}

func getBranch() string {
	return branchPrefix + "_" + time.Now().Format("2006-01-02T15-04-05")
}

func getOrigin(githubToken string) string {
	return strings.Replace(repoURL, "github.com", gitUser+":"+githubToken+"@github.com", -1)
}

func strPtr(s string) *string {
	return &s
}

func newTrue() *bool {
	b := true
	return &b
}
