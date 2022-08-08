package main

import (
	"errors"
	"os"
)

func LoadCredentialsFromEnvironment() (login string, password string, err error) {
	login = os.Getenv("login")
	password = os.Getenv("password")

	if login != "" && password != "" {
		return login, password, nil
	}

	return "", "", errors.New("Error loading credentials from environment. (login, password)")
}
