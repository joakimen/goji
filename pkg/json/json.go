package json

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ToJSON[T any](data T) (string, error) {
	encoded, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error encoding JSON: %w", err)
	}
	return string(encoded), nil
}

func FromJSON[T any](data string) (T, error) {
	var t T
	err := json.Unmarshal([]byte(data), &t)
	return t, err
}

func Format(data string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(data), "", "  ")
	if err != nil {
		return "", fmt.Errorf("error pretty printing JSON: %w", err)
	}
	return prettyJSON.String(), nil
}
