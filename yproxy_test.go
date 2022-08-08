package main

import (
	"os"
	"testing"
)

func TestHappyPath(t *testing.T) {

	username := os.Getenv("GLOVEBOX_USER")
	password := os.Getenv("GLOVEBOX_PASS")

	yproxy := yproxy{}
	err := yproxy.Login(username, password)

	if err != nil {
		t.Errorf("No auth fail.")
	}

	page := yproxy.browser.MustPage("https://news.ycombinator.com/newest")
	el := page.MustElement("#me")

	if el.MustText() != "rcountry21" {
		t.Errorf("Cannot find correct user after login.")
	}
}

func TestBadPassword(t *testing.T) {

	username := "GLOVEBOX_USER"
	password := "GLOVEBOX_PASS"

	yproxy := yproxy{}
	err := yproxy.Login(username, password)

	if err == nil {
		t.Errorf("Auth should fail.")
	}

}
