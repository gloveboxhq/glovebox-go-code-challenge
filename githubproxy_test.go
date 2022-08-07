package main

import (
	"testing"
)

func TestGithubProxyHasLogin(t *testing.T) {
	var sut PolicyProvider = GitHubProvider{}
	sut.Login("", "")
}
