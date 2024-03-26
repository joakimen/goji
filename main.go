package main

import (
	"fmt"
	"github.com/joakimen/goji/pkg/jira"
	"log"
	"os"
)

func main() {

	apiCredentials := jira.APICredentials{
		User:  RequireEnv("JIRA_API_USER"),
		Token: RequireEnv("JIRA_API_TOKEN"),
		Host:  RequireEnv("JIRA_HOST"),
	}

	issues, err := jira.Search(apiCredentials)
	if err != nil {
		log.Fatalf("error fetching jira issues: %v", err)
	}
	fmt.Println(issues)
}

func RequireEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}
	return value
}
