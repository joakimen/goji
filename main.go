package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type JiraCreds struct {
	User  string
	Token string
	Host  string
}

type Issue struct {
	Key    string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
		Created string `json:"created"`
		Status  struct {
			Name string `json:"name"`
		} `json:"status"`
	} `json:"fields"`
}

type Issues []Issue

type JiraSearchResults struct {
	Issues          Issues   `json:"issues"`
	MaxResults      int      `json:"maxResults"`
	StartAt         int      `json:"startAt"`
	Total           int      `json:"total"`
	WarningMessages []string `json:"warningMessages"`
}

func (issues Issues) String() string {
	issuesJSON, err := json.MarshalIndent(issues, "", "  ")
	if err != nil {
		return fmt.Sprintf("error marshalling issues to json: %v", err)
	}
	return string(issuesJSON)
}

func main() {

	jira := JiraCreds{
		User:  RequireEnv("JIRA_API_USER"),
		Token: RequireEnv("JIRA_API_TOKEN"),
		Host:  RequireEnv("JIRA_HOST"),
	}

	var searchResults JiraSearchResults
	err := SearchJiraIssues(jira, &searchResults)
	if err != nil {
		log.Fatalf("error fetching jira issues: %v", err)
	}
	fmt.Println(searchResults.Issues)
}

func SearchJiraIssues(jira JiraCreds, j *JiraSearchResults) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest("GET", jira.Host+"/rest/api/3/search", nil)
	req.SetBasicAuth(jira.User, jira.Token)

	queryParams := req.URL.Query()
	queryParams.Add("jql", "assignee = currentUser() AND resolution = Unresolved")
	queryParams.Add("fields", "summary,status,created")
	req.URL.RawQuery = queryParams.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error fetching jira issues: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("error closing response body: %v", err)
		}
	}(resp.Body)

	err = json.Unmarshal(body, j)
	if err != nil {
		return err
	}
	return nil
}

func RequireEnv(envVar string) string {
	val := os.Getenv(envVar)
	if val == "" {
		log.Fatalf("missing required environment variable %s", envVar)
	}
	return val
}
