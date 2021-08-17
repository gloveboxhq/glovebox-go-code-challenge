package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGithubCrawler(t *testing.T) {
	godotenv.Load()
	providerMeta := ProviderData{"github", os.Getenv("github"), []Route{}}
	requestMeta := AccessRequest{"test", "test"}
	crawler(providerMeta, requestMeta)
}

func TestHumbleBundleCrawler(t *testing.T) {
	godotenv.Load()
	providerMeta := ProviderData{"humbleBundle", os.Getenv("humbleBundle"), []Route{}}
	requestMeta := AccessRequest{"test", "test"}
	crawler(providerMeta, requestMeta)
}
