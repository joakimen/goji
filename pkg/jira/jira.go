package jira

import (
	"context"
	"fmt"
	"time"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
	"github.com/joakimen/goji/pkg/auth"
)

type Epic struct {
	Key     string
	Summary string
	Created time.Time
	Status  string
}

func NewEpic(issue jira.Issue) Epic {
	return Epic{
		Key:     issue.Key,
		Summary: issue.Fields.Summary,
		Created: time.Time(issue.Fields.Created),
		Status:  issue.Fields.Status.Name,
	}
}

func NewClient(creds auth.Credentials) (*jira.Client, error) {
	jiraURL := creds.Host
	tp := jira.BasicAuthTransport{
		Username: creds.Username,
		APIToken: creds.APIToken,
	}

	client, err := jira.NewClient(jiraURL, tp.Client())
	if err != nil {
		panic(err)
	}

	return client, nil
}

func ListEpics(jiraClient *jira.Client, projectID string, all bool, mine bool) ([]Epic, error) {
	jql := fmt.Sprintf("project = %s AND issuetype = Epic", projectID)
	if !all {
		jql += " AND resolution = Unresolved"
	}

	if mine {
		jql += " AND assignee = currentUser()"
	}

	opts := &jira.SearchOptions{
		StartAt:    0,
		MaxResults: 50,
		Fields:     []string{"key", "summary", "created", "status"},
	}

	var epics []Epic
	for {
		issues, resp, err := jiraClient.Issue.Search(context.Background(), jql, opts)
		if err != nil {
			return nil, fmt.Errorf("failed when searching for epics for projectID '%s': %w", projectID, err)
		}
		for _, issue := range issues {
			epics = append(epics, NewEpic(issue))
		}
		if resp.StartAt+resp.MaxResults >= resp.Total {
			break
		}
		// Update StartAt for the next batch
		opts.StartAt += opts.MaxResults
	}
	return epics, nil
}

type Issue struct {
	Key     string
	Summary string
	Created time.Time
	Status  string
	Type    string
}

func NewIssue(issue jira.Issue) Issue {
	return Issue{
		Key:     issue.Key,
		Summary: issue.Fields.Summary,
		Created: time.Time(issue.Fields.Created),
		Status:  issue.Fields.Status.Name,
		Type:    issue.Fields.Type.Name,
	}
}

func ListBugs(jiraClient *jira.Client, projectID string, all bool, mine bool, limit int) ([]Issue, error) {
	jql := fmt.Sprintf("project = %s AND issuetype = Bug", projectID)
	if !all {
		jql += " AND resolution = Unresolved"
	}

	if mine {
		jql += " AND assignee = currentUser()"
	}

	opts := &jira.SearchOptions{
		StartAt:    0,
		MaxResults: limit,
		Fields:     []string{"key", "summary", "created", "status", "issuetype"},
	}

	var bugs []Issue
	for {
		searchResults, resp, err := jiraClient.Issue.Search(context.Background(), jql, opts)
		if err != nil {
			return nil, fmt.Errorf("failed when searching for bugs for projectID '%s': %w", projectID, err)
		}
		for _, issue := range searchResults {
			bugs = append(bugs, NewIssue(issue))
		}
		if resp.StartAt+resp.MaxResults >= resp.Total || len(bugs) >= limit {
			break
		}
		opts.StartAt += opts.MaxResults
	}

	if len(bugs) > limit {
		bugs = bugs[:limit]
	}
	return bugs, nil
}

func ListIssues(jiraClient *jira.Client, projectID string, all bool, mine bool, limit int) ([]Issue, error) {
	jql := fmt.Sprintf("project = %s AND issuetype != Epic", projectID)
	if !all {
		jql += " AND resolution = Unresolved"
	}

	if mine {
		jql += " AND assignee = currentUser()"
	}

	opts := &jira.SearchOptions{
		StartAt:    0,
		MaxResults: limit,
		Fields:     []string{"key", "summary", "created", "status", "issuetype"},
	}

	var issues []Issue
	for {
		searchResults, resp, err := jiraClient.Issue.Search(context.Background(), jql, opts)
		if err != nil {
			return nil, fmt.Errorf("failed when searching for issues for projectID '%s': %w", projectID, err)
		}
		for _, issue := range searchResults {
			issues = append(issues, NewIssue(issue))
		}
		if resp.StartAt+resp.MaxResults >= resp.Total || len(issues) >= limit {
			break
		}
		opts.StartAt += opts.MaxResults
	}

	if len(issues) > limit {
		issues = issues[:limit]
	}
	return issues, nil
}
