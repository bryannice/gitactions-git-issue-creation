package configuration

import (
	"errors"
	"os"
)

type Config struct {
	GitHubToken      string
	GitHubSha        string
	GitHubOwner      string
	GitHubRepository string
	GitHubIssueTitle string
	GitHubIssueBody  string
}

func (config *Config) Init() error {

	var err error

	// Checking and setting required environment variables
	if os.Getenv("GITHUB_TOKEN") == "" {
		err = errors.New("GITHUB_TOKEN must be set")
	} else {
		config.GitHubToken = os.Getenv("GITHUB_TOKEN")
	}

	if os.Getenv("GITHUB_SHA") == "" {
		err = errors.New("GITHUB_SHA must be set")
	} else {
		config.GitHubSha = os.Getenv("GITHUB_SHA")
	}

	if os.Getenv("GITHUB_OWNER") == "" {
		err = errors.New("GITHUB_OWNER must be set")
	} else {
		config.GitHubOwner = os.Getenv("GITHUB_OWNER")
	}

	if os.Getenv("GITHUB_REPOSITORY") == "" {
		err = errors.New("GITHUB_REPOSITORY must be set")
	} else {
		config.GitHubRepository = os.Getenv("GITHUB_REPOSITORY")
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
