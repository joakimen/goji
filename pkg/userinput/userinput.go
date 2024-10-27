package userinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func ReadString(prompt string) (string, error) {
	fmt.Print(prompt)
	r := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	userInput := scanner.Text()
	return strings.TrimSpace(userInput), nil
}

func ReadStringMasked(prompt string) (string, error) {
	fmt.Print(prompt)
	fileDescriptor := int(os.Stdin.Fd())
	byteVal, err := term.ReadPassword(fileDescriptor)
	fmt.Println()
	if err != nil {
		return "", fmt.Errorf("error reading password: %w", err)
	}
	maskedString := string(byteVal)
	return strings.TrimSpace(maskedString), nil
}
