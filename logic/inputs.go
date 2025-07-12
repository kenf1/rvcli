package logic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// confirm fields to create jwt token present
func CheckInputs(userConfig UserConfig, apiConfig ApiConfig) error {
	var messages []string

	if userConfig.Username == "" {
		messages = append(messages, "Username is missing")
	}
	if userConfig.Password == "" {
		messages = append(messages, "Password is missing")
	}
	if userConfig.Fullname == "" {
		messages = append(messages, "Fullname is missing")
	}
	if userConfig.Email == "" {
		messages = append(messages, "Email is missing")
	}
	if apiConfig.Host == "" {
		messages = append(messages, "Host is missing")
	}

	if len(messages) > 0 {
		return fmt.Errorf("%s", strings.Join(messages, "; "))
	}

	return nil
}

// prompt text input
func promptText(prompt string) (string, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	return strings.TrimSpace(input), nil
}

// prompt text wrapper
func PromptTextWrapper(input []string) error {
	for _, item := range input {
		value, err := promptText(fmt.Sprintf("Enter %s: ", item))
		if err != nil {
			return fmt.Errorf("Invalid entry for %s", item)
		}

		os.Setenv(strings.ToUpper(item), value)
	}
	return nil
}
