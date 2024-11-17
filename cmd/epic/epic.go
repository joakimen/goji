package epic

import (
	"fmt"
	"os"

	"github.com/joakimen/goji/pkg/auth"
	"github.com/joakimen/goji/pkg/format"
	"github.com/joakimen/goji/pkg/jira"
	"github.com/joakimen/goji/pkg/json"
)

func List(projectID string, jsonOutput bool, all bool, mine bool) error {
	stderr := os.Stderr
	fmt.Fprintln(stderr, "Listing epics")
	creds, err := auth.GetCredentials()
	if err != nil {
		return fmt.Errorf("failed to get credentials: %w", err)
	}
	jiraClient, err := jira.NewClient(creds)
	if err != nil {
		return fmt.Errorf("failed to create Jira client: %w", err)
	}

	epics, err := jira.ListEpics(jiraClient, projectID, all, mine)
	if err != nil {
		return fmt.Errorf("failed to list epics: %w", err)
	}

	fmt.Fprintf(stderr, "Listing all %d returned epics\n", len(epics))
	if jsonOutput {
		epicsJson, err := json.ToJSON(epics)
		if err != nil {
			return fmt.Errorf("failed to encode epics to JSON: %w", err)
		}
		epicsJsonPretty, err := json.Format(epicsJson)
		if err != nil {
			return fmt.Errorf("failed to format JSON: %w", err)
		}
		fmt.Println(epicsJsonPretty)
	} else {
		for _, epic := range epics {
			fmt.Println(format.FormatItem(
				epic.Key,
				epic.Status,
				epic.Summary,
				epic.Created,
			))
		}
	}
	fmt.Fprintln(stderr, "Done listing epics")
	return nil
}
