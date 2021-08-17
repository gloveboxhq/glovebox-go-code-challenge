package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGithubProviderLogin(t *testing.T) {
	godotenv.Load()
	// manually creating a provider to limits scope of code in test
	provider := ProviderData{"github", os.Getenv("github"), []Route{}}
	provider.Login(os.Getenv("username"), os.Getenv("password"))
}

func TestHumbleBundleProviderLogin(t *testing.T) {
	godotenv.Load()
	// manually creating a provider to limits scope of code in test
	provider := ProviderData{"humbleBundle", os.Getenv("humbleBundle"), []Route{}}
	provider.Login(os.Getenv("username"), os.Getenv("password"))
}
