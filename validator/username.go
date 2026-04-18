package validator

import (
	"errors"
	"regexp"
	"strings"
)

var usernameRegex = regexp.MustCompile(`^[A-Za-z0-9-]+$`)

func ValidateGitHubUsername(s string) error {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return errors.New("username cannot be empty")
	}

	if strings.Contains(s, " ") {
		return errors.New("username cannot contain spaces")
	}

	if len(s) > 39 {
		return errors.New("username must be at most 39 characters")
	}

	if strings.HasPrefix(s, "-") || strings.HasSuffix(s, "-") {
		return errors.New("username cannot start or end with a hyphen")
	}

	if strings.Contains(s, "--") {
		return errors.New("username cannot contain consecutive hyphens")
	}

	if !usernameRegex.MatchString(s) {
		return errors.New("username can only contain letters, numbers, and hyphens")
	}

	return nil
}
