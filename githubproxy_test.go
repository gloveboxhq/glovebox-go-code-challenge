package main

import (
	"os"
	"testing"
)

func TestGithubProxyHasLogin(t *testing.T) {
	var sut PolicyProvider = GitHubProvider{}
	sut.Login("", "")
}

func TestCredentialsLoadedFromEnvironment(t *testing.T) {
	const mock_username = "MOCK_USERNAME"
	const mock_password = "MOCK_PASSWORD"
	os.Setenv("username", mock_username)
	os.Setenv("password", mock_password)

	username, password, credentialsError := LoadCredentialsFromEnvironment()

	if credentialsError != nil {
		t.Fatalf("Method returned error")
	}

	if username != mock_username || password != mock_password {
		t.Fatalf("Username or Password env vars returned unexpected value")
	}
}
