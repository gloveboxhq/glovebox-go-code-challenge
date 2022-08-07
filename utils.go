package main

import (
	"errors"
	"os"
)

func LoadCredentialsFromEnvironment() (username string, password string, err error) {
	username = os.Getenv("username")
	password = os.Getenv("password")

	if username != "" && password != "" {
		return username, password, nil
	}

	return "", "", errors.New("Error loading credentials from environment. (username, password)")
}
