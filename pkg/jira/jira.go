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
