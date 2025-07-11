package logic

import (
	"fmt"
	"strings"
)

func CheckInputs(userConfig UserConfig) error {
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

	if len(messages) > 0 {
		return fmt.Errorf("%s", strings.Join(messages, "; "))
	}

	return nil
}
