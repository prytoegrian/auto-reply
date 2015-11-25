package common

import (
	"log"
	"os"

	"github.com/parkr/auto-reply/Godeps/_workspace/src/github.com/google/go-github/github"
	"github.com/parkr/auto-reply/Godeps/_workspace/src/golang.org/x/oauth2"
)

const accessTokenEnvVar = "GITHUB_ACCESS_TOKEN"

func GitHubToken() string {
	return os.Getenv(accessTokenEnvVar)
}

func NewClient() *github.Client {
	if token := GitHubToken(); token != "" {
		return github.NewClient(oauth2.NewClient(
			oauth2.NoContext,
			oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
			),
		))
	} else {
		log.Fatalf("%s required", accessTokenEnvVar)
		return nil
	}
}