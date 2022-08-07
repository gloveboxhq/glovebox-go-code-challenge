package githubproxy

import (
	"errors"
	"os"
)

type GitHubProvider struct {
}

func (ghp GitHubProvider) Login(username, password string) error {
	return nil
}

func LoadCredentialsFromEnvironment() (username string, password string, err error) {
	username = os.Getenv("username")
	password = os.Getenv("password")

	if username != "" && password != "" {
		return username, password, nil
	}

	return "", "", errors.New("Error loading credentials from environment. (username, password)")
}
