package configuration

import (
	"errors"
	"os"
)

type Config struct {
	GitHubToken     string
	GitHubCommitSha       string
	GitHubRepoOwner     string
	GitHubRepoName     string
	GitHubIssueTitle     string
	GitHubIssueBody       string
}

func (config *Config) Init() error {

	var err error

	// Checking and setting required environment variables
	if os.Getenv("GITHUB_TOKEN") == "" {
		err = errors.New("GITHUB_TOKEN must be set")
	} else {
		config.GitHubToken = os.Getenv("GITHUB_TOKEN")
	}

	if os.Getenv("GITHUB_COMMIT_SHA") == "" {
		err = errors.New("GITHUB_COMMIT_SHA must be set")
	} else {
		config.GitHubCommitSha = os.Getenv("GITHUB_COMMIT_SHA")
	}

	if os.Getenv("GITHUB_REPO_OWNER") == "" {
		err = errors.New("GITHUB_REPO_OWNER must be set")
	} else {
		config.GitHubRepoOwner = os.Getenv("GITHUB_REPO_OWNER")
	}

	if os.Getenv("GITHUB_REPO_NAME") == "" {
		err = errors.New("GITHUB_REPO_NAME must be set")
	} else {
		config.GitHubRepoName = os.Getenv("GITHUB_REPO_NAME")
	}

	if os.Getenv("GITHUB_ISSUE_TITLE") == "" {
		err = errors.New("GITHUB_ISSUE_TITLE must be set")
	} else {
		config.GitHubIssueTitle = os.Getenv("GITHUB_ISSUE_TITLE")
	}

	if os.Getenv("GITHUB_ISSUE_BODY") == "" {
		err = errors.New("GITHUB_ISSUE_BODY must be set")
	} else {
		config.GitHubIssueBody = os.Getenv("GITHUB_ISSUE_BODY")
	}

	return err
}
