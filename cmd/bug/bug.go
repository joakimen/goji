package bug

import (
	"fmt"
	"os"

	"github.com/joakimen/goji/pkg/auth"
	"github.com/joakimen/goji/pkg/format"
	"github.com/joakimen/goji/pkg/jira"
	"github.com/joakimen/goji/pkg/json"
)

func List(projectID string, jsonOutput bool, all bool, mine bool, limit int) error {
	stderr := os.Stderr
	fmt.Fprintln(stderr, "Listing bugs")

	creds, err := auth.GetCredentials()
	if err != nil {
		return fmt.Errorf("failed to get credentials: %w", err)
	}

	jiraClient, err := jira.NewClient(creds)
	if err != nil {
		return fmt.Errorf("failed to create Jira client: %w", err)
	}

	bugs, err := jira.ListBugs(jiraClient, projectID, all, mine, limit)
	if err != nil {
		return fmt.Errorf("failed to list bugs: %w", err)
	}

	fmt.Fprintf(stderr, "Listing %d returned bugs\n", len(bugs))
	if jsonOutput {
		bugsJson, err := json.ToJSON(bugs)
		if err != nil {
			return fmt.Errorf("failed to encode bugs to JSON: %w", err)
		}
		bugsJsonPretty, err := json.Format(bugsJson)
		if err != nil {
			return fmt.Errorf("failed to format JSON: %w", err)
		}
		fmt.Println(bugsJsonPretty)
	} else {
		for _, bug := range bugs {
			fmt.Println(format.FormatItem(
				bug.Key,
				bug.Status,
				bug.Summary,
				bug.Created,
			))
		}
	}
	fmt.Fprintln(stderr, "Done listing bugs")
	return nil
}
