package auth

import (
	"context"
	"fmt"

	"github.com/joakimen/goji/pkg/auth"
	"github.com/joakimen/goji/pkg/jira"
	"github.com/joakimen/goji/pkg/userinput"
)

func Login() error {
	fmt.Println("Enter your Jira credentials:")
	username, err := userinput.ReadString("Username: ")
	if err != nil {
		return fmt.Errorf("failed to read username: %w", err)
	}

	apiToken, err := userinput.ReadStringMasked("APIToken: ")
	if err != nil {
		return fmt.Errorf("failed to read password: %w", err)
	}

	host, err := userinput.ReadString("Jira host URL: ")
	if err != nil {
		return fmt.Errorf("failed to read Jira host URL: %w", err)
	}

	jiraCreds := auth.Credentials{
		Username: username,
		APIToken: apiToken,
		Host:     host,
	}
	fmt.Println("Storing credentials in keyring...")
	err = auth.SetCredentials(jiraCreds)
	if err != nil {
		return fmt.Errorf("failed to store credentials: %w", err)
	}
	fmt.Println("Credentials stored.")
	return nil
}

func Show() error {
	storedCreds, err := auth.GetCredentials()
	if err != nil {
		return fmt.Errorf("failed to get credentials: %w", err)
	}
	fmt.Println(storedCreds)
	return nil
}

func Status() error {
	creds, err := auth.GetCredentials()
	if err != nil {
		fmt.Println("❌ Not authenticated - no credentials found")
		return nil
	}

	// Create a Jira client and test the credentials
	client, err := jira.NewClient(creds)
	if err != nil {
		fmt.Printf("❌ Authentication failed: %v\n", err)
		return nil
	}

	// Verify credentials by making a test API call
	_, _, err = client.User.GetCurrentUser(context.Background())
	if err != nil {
		fmt.Printf("❌ Authentication failed for user %s: %v\n", creds.Username, err)
		return nil
	}

	fmt.Printf("✅ Authenticated as %s\n", creds.Username)
	return nil
}
