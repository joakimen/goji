package issue

import (
	"fmt"
	"os"

	"github.com/joakimen/goji/pkg/auth"
	"github.com/joakimen/goji/pkg/jira"
	"github.com/joakimen/goji/pkg/json"
)

func List(projectID string, jsonOutput bool, all bool, mine bool, limit int) error {
	stderr := os.Stderr
	fmt.Fprintln(stderr, "Listing issues")

	creds, err := auth.GetCredentials()
	if err != nil {
		return fmt.Errorf("failed to get credentials: %w", err)
	}

	jiraClient, err := jira.NewClient(creds)
	if err != nil {
		return fmt.Errorf("failed to create Jira client: %w", err)
	}

	issues, err := jira.ListIssues(jiraClient, projectID, all, mine, limit)
	if err != nil {
		return fmt.Errorf("failed to list issues: %w", err)
	}

	fmt.Fprintf(stderr, "Listing %d returned issues\n", len(issues))
	if jsonOutput {
		issuesJson, err := json.ToJSON(issues)
		if err != nil {
			return fmt.Errorf("failed to encode issues to JSON: %w", err)
		}
		issuesJsonPretty, err := json.Format(issuesJson)
		if err != nil {
			return fmt.Errorf("failed to format JSON: %w", err)
		}
		fmt.Println(issuesJsonPretty)
	} else {
		for _, issue := range issues {
			fmt.Printf("[%s] [%s] %s (%s)\n", issue.Key, issue.Status, issue.Summary, issue.Created)
		}
	}
	fmt.Fprintln(stderr, "Done listing issues")
	return nil
}
