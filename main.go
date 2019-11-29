package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bryan-nice/git-issue-creation/configuration"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"log"
)

func main() {
	var issueRequest *github.IssueRequest
	var result []byte
	var err error
	var config *configuration.Config

	// Init configuration from environment variables
	config = new(configuration.Config)
	err = config.Init()
	if err != nil {
		log.Fatalf("Exception: %v", err)
	}

	// Required variables to create an Issue
	gitHubToken := config.GitHubToken
	gitHubSha := config.GitHubSha
	gitHubOwner := config.GitHubOwner
	gitHubRepository := config.GitHubRepository
	gitHubIssueTitle := config.GitHubIssueTitle
	gitHubIssueBody := config.GitHubIssueBody

	// Search string
	searchString := fmt.Sprintf("repo:%s is:issue is:open %s", fmt.Sprintf("%s/%s", gitHubOwner, gitHubRepository), gitHubSha)

	// Authenticating and creating GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitHubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Search issues to check if it exists already for a specific commit SHA
	searchResults, _, err := client.Search.Issues(ctx, searchString, nil)
	if err != nil {
		log.Printf("%+v", errors.Wrap(err, "Exception"))
	}

	if len(searchResults.Issues) == 0 {
		// Create issue if not exist for the commit SHA
		issueRequest = new(github.IssueRequest)
		issueRequest.Title = &gitHubIssueTitle
		issueRequest.Body = &gitHubIssueBody
		issueCreated, _, err := client.Issues.Create(ctx, gitHubOwner, gitHubRepository, issueRequest)
		if err != nil {
			log.Printf("%+v", errors.Wrap(err, "Exception"))
		}
		//fmt.Printf("%+v", issueCreated)
		result, err = json.Marshal(issueCreated.HTMLURL)
		if err != nil {
			log.Printf("%+v", errors.Wrap(err, "Exception"))
		}
	} else {
		// Return issue if it exists
		result, err = json.Marshal(searchResults.Issues[0].HTMLURL)
		if err != nil {
			log.Printf("%+v", errors.Wrap(err, "Exception"))
		}
	}

	fmt.Printf(fmt.Sprintf("::set-output name=git_issue_url::%s", string(result)))
}
