package main

import (
	"testing"
)

func TestGithubProxyHasLogin(t *testing.T) {
	var sut PolicyProvider = GitHubProvider{}
	sut.Login("", "")
}

func TestGithubProxyLogsInWithEnvironmentVars(t *testing.T) {
	var sut PolicyProvider = GitHubProvider{}
	login, password, environmentVarErrror := LoadCredentialsFromEnvironment()

	if environmentVarErrror != nil {
		t.Fatalf("Failure loading environment variables")
	}

	loginError := sut.Login(login, password)

	if loginError != nil {
		t.Fatalf("Login failed")
	}
}
