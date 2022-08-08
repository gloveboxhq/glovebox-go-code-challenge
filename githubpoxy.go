package main

import (
	"errors"
	"io"
	"time"

	"github.com/go-rod/rod"
)

type GitHubProvider struct {
}

func (ghp GitHubProvider) Login(login, password string) error {
	const loginURL = "https://github.com/login"

	// Open page
	browser := rod.New()
	err := browser.Connect()

	if err != nil {
		return errors.New("Unable to connect to broswer")
	}

	defer browser.MustClose()
	browser.ControlURL(loginURL)
	err = browser.Connect()
	if err != nil {
		return errors.New("Unable to open GitHub login URL")
	}
	page := browser.MustPage(loginURL)

	// Enter credentials
	loginElement, err := page.Element("#login_field")

	if err != nil {
		return errors.New("Unable to find login text field")
	}

	passwordElement, err := page.Element("#password")

	if err != nil {
		return errors.New("Unable to find password text field")
	}

	loginElement.MustClick().MustInput(login)
	passwordElement.MustClick().MustInput(password)

	// Click signin button
	// Login button did not have its own unique ID.  Had to use a complex
	// 	value taken from Chrome dev tools.
	signInButtonElement, err := page.Element("#login > div.auth-form-body.mt-3 > form > div > input.btn.btn-primary.btn-block.js-sign-in-button")
	if err != nil {
		return errors.New("Unable to find sign in button")
	}
	signInButtonElement.MustClick()
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
