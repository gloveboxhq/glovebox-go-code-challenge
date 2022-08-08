package main

import (
	"io"
	"time"

	"github.com/go-rod/rod"
)

type GitHubProvider struct {
}

func (ghp GitHubProvider) Login(login, password string) error {
	// Open page
	browser := rod.New().MustConnect()
	defer browser.MustClose()
	page := browser.MustPage("https://github.com/login")

	// Enter credentials
	page.MustElement("#login_field").MustClick().MustInput(login)
	page.MustElement("#password").MustClick().MustInput(password)

	// Click signin button
	// Login button did not have its own unique ID.  Had to use a complex
	// 	value taken from Chrome dev tools.
	page.MustElement("#login > div.auth-form-body.mt-3 > form > div > input.btn.btn-primary.btn-block.js-sign-in-button").MustClick()
	page.WaitLoad()

	// Screenshot to verify
	filename := time.Now().Format("20060102150405") + "_go_test.png"
	page.MustScreenshot(filename)

	return nil
}

func (ghp GitHubProvider) Policies() ([]Policy, error) {
	return nil, nil
}

func (ghp GitHubProvider) DocumentDownload(downloadKey string) (io.ReadCloser, error) {
	return nil, nil
}
