package main

import (
	"os"
	"testing"
)

func TestCredentialsLoadedFromEnvironment(t *testing.T) {
	const mock_login = "MOCK_LOGIN"
	const mock_password = "MOCK_PASSWORD"
	os.Setenv("login", mock_login)
	os.Setenv("password", mock_password)

	login, password, credentialsError := LoadCredentialsFromEnvironment()

	if credentialsError != nil {
		t.Fatalf("Method returned error")
	}

	if login != mock_login || password != mock_password {
		t.Fatalf("Login or Password env vars returned unexpected value")
	}
}
