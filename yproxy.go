package main

import (
	"errors"
	"fmt"

	"github.com/go-rod/rod"
)

type yproxy struct {
	loginUrl            string
	usernameSelector    string
	passwordSelector    string
	loginButtonSelector string
	browser             *rod.Browser
}

// Implementing Login method with PolicyProvider interface
func (proxy *yproxy) Login(username, password string) error {
	// Setting selector data to find webpage and login/button click info
	proxy.loginUrl = "https://news.ycombinator.com/login?goto=newest"
	proxy.usernameSelector = "body > form:nth-child(4) > table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(2) > input:nth-child(1)"
	proxy.passwordSelector = "body > form:nth-child(4) > table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(2) > input:nth-child(1)"
	proxy.loginButtonSelector = "body > form:nth-child(4) > input[type=submit]:nth-child(4)"

	// Connect to login url
	proxy.browser = rod.New().MustConnect().NoDefaultDevice()
	page := proxy.browser.MustPage(proxy.loginUrl)

	// try to login and return any errors encountered
	err := rod.Try(func() {
		page.MustElement(proxy.usernameSelector).MustInput(username)
		page.MustElement(proxy.passwordSelector).MustInput(password)
		page.MustElement(proxy.loginButtonSelector).MustClick()
	})
	if err != nil {
		fmt.Println("return err")
		return err
	}

	// verify login was successful
	page = proxy.browser.MustPage("https://news.ycombinator.com/newest")
	el, err := page.Element("#me")
	fmt.Println(el.MustText())
	if el != nil && el.MustText() == username {
		return nil
	} else {
		return errors.New("Login failed: " + err.Error())
	}
}
