package main

import (
	"testing"
)

func TestGithubProxyHasLogin(t *testing.T) {
	var sut PolicyProvider = GitHubProvider{}
	sut.Login("", "")
}

func TestCredentialsLoadedFromEnvironment(t *testing.T) {
	username, password, credentialsError := LoadCredentialsFromEnvironment()

	if credentialsError != nil {
		t.Fatalf("Method returned error")
	}

	if username == "" || password == "" {
		t.Fatalf("Username or Password returned empty string")
	}
}
