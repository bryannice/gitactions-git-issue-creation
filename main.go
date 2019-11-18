package main

import (
	"github.com/bryan-nice/git-issue-creation/configuration"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"context"
	"fmt"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
)

func main(){
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
	gitHubCommitSha := config.GitHubCommitSha
	gitHubRepoOwner := config.GitHubRepoOwner
	gitHubRepoName := config.GitHubRepoName
	gitHubIssueTitle := config.GitHubIssueTitle
	gitHubIssueBody := config.GitHubIssueBody

	// Search string
	searchString := fmt.Sprintf("repo:%s is:issue is:open %s", fmt.Sprintf("%s/%s",gitHubRepoOwner,gitHubRepoName), gitHubCommitSha)

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
		issueRequest.Title = &issueTitle
		issueRequest.Body = &issueBody
		issueCreated, _, err := client.Issues.Create(ctx, gitHubRepoOwner, gitHubRepository, issueRequest)
		if err != nil {
			log.Printf("%+v", errors.Wrap(err, "Exception"))
		}
		result, err = json.Marshal(issueCreated)
		if err != nil {
			log.Printf("%+v", errors.Wrap(err, "Exception"))
		}
	} else {
		// Return issue if it exists
		result, err = json.Marshal(searchResults.Issues)
		if err != nil {
			log.Printf("%+v", errors.Wrap(err, "Exception"))
		}
	}

	fmt.Printf(string(result))
}
