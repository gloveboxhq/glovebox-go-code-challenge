package main

import (
	"errors"
	"io"
	"os"
)

type GitHubProvider struct {
}

func (ghp GitHubProvider) Login(username, password string) error {
	return nil
}

func (ghp GitHubProvider) Policies() ([]Policy, error) {
	return nil, nil
}

func (ghp GitHubProvider) DocumentDownload(downloadKey string) (io.ReadCloser, error) {
	return nil, nil
}
