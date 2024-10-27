package auth

import (
	"fmt"

	"github.com/joakimen/goji/pkg/json"
	"github.com/zalando/go-keyring"
)

type Credentials struct {
	Username string
	APIToken string
	Host     string
}

type KeyringConfiguration struct {
	Service string
	Item    string
}

const (
	KeyringService = "goji"
	KeyringItem    = "jira-creds"
)

var KeyringCfg = KeyringConfiguration{
	Service: KeyringService,
	Item:    KeyringItem,
}

func (creds Credentials) String() string {
	credsJson, err := ToJSON(creds)
	if err != nil {
		return fmt.Sprintf("error encoding credentials json: %s", err)
	}

	credsJsonPretty, err := json.Format(credsJson)
	if err != nil {
		return fmt.Sprintf("error formatting credentials json: %s", err)
	}
	return credsJsonPretty
}

func ToJSON(creds Credentials) (string, error) {
	credsJson, err := json.ToJSON(creds)
	if err != nil {
		return "", fmt.Errorf("error encoding credentials: %w", err)
	}
	return credsJson, nil
}

func ToCredentials(data string) (Credentials, error) {
	creds, err := json.FromJSON[Credentials](data)
	if err != nil {
		return Credentials{}, fmt.Errorf("error decoding credentials: %w", err)
	}
	return creds, nil
}

func GetCredentials() (Credentials, error) {
	credsJson, err := keyring.Get(KeyringCfg.Service, KeyringCfg.Item)
	if err != nil {
		return Credentials{}, fmt.Errorf("failed to get keyring item '%+v': %w", KeyringCfg, err)
	}
	return ToCredentials(credsJson)
}

func SetCredentials(creds Credentials) error {
	credsJson, err := ToJSON(creds)
	if err != nil {
		return fmt.Errorf("failed to encode credentials to JSON: %w", err)
	}
	err = keyring.Set(KeyringCfg.Service, KeyringCfg.Item, credsJson)
	if err != nil {
		return fmt.Errorf("failed to save credentials to keyring: %w", err)
	}
	return nil
}
