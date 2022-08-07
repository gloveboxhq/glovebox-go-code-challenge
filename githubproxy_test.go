package main

import (
	"testing"
)

func TestGithubProxyHasLogin(t *testing.T) {
	var provider PolicyProvider = GitHubProvider{}
	provider.Login("", "")
}
