package jira

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// APICredentials required to authenticate with the Jira API.
type APICredentials struct {
	User  string
	Token string
	Host  string
}

// Issue represents a Jira issue.
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

// SearchResults represents the response from a Jira Issue Search.
type SearchResults struct {
	Issues          Issues   `json:"issues"`
	MaxResults      int      `json:"maxResults"`
	StartAt         int      `json:"startAt"`
	Total           int      `json:"total"`
	WarningMessages []string `json:"warningMessages"`
}

// String returns a pretty-printed JSON representation of the issues.
func (issues Issues) String() string {
	issuesJSON, err := json.MarshalIndent(issues, "", "  ")
	if err != nil {
		return fmt.Sprintf("error marshalling issues to json: %v", err)
	}
	return string(issuesJSON)
}

func httpGet(client http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing request: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("error closing response body: %v", err)
		}
	}(resp.Body)
	return body, nil
}

// Search for unresolved Jira issues assigned to the current user.
func Search(jira APICredentials) (Issues, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest("GET", jira.Host+"/rest/api/3/search", nil)
	req.SetBasicAuth(jira.User, jira.Token)

	queryParams := req.URL.Query()
	queryParams.Add("jql", "assignee = currentUser() AND resolution = Unresolved")
	queryParams.Add("fields", "summary,status,created")
	req.URL.RawQuery = queryParams.Encode()

	body, err := httpGet(*client, req)
	if err != nil {
		return nil, fmt.Errorf("error performing http get: %w", err)
	}

	var searchResults SearchResults
	err = json.Unmarshal(body, &searchResults)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return searchResults.Issues, nil
}
