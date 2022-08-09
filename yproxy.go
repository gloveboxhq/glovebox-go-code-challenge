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

func (proxy *yproxy) Login(username, password string) error {
	proxy.loginUrl = "https://news.ycombinator.com/login?goto=newest"
	proxy.usernameSelector = "body > form:nth-child(4) > table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(2) > input:nth-child(1)"
	proxy.passwordSelector = "body > form:nth-child(4) > table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(2) > input:nth-child(1)"
	proxy.loginButtonSelector = "body > form:nth-child(4) > input[type=submit]:nth-child(4)"

	proxy.browser = rod.New().MustConnect().NoDefaultDevice()
	page := proxy.browser.MustPage(proxy.loginUrl)

	err := rod.Try(func() {
		page.MustElement(proxy.usernameSelector).MustInput(username)
		page.MustElement(proxy.passwordSelector).MustInput(password)
		page.MustElement(proxy.loginButtonSelector).MustClick()
	})
	if err != nil {
		fmt.Println("return err")
		return err
	}

	el, err := page.Element("#me")
	fmt.Println(el.MustText())
	if el != nil && el.MustText() == username {
		return nil
	} else {
		return errors.New("Login failed: " + err.Error())
	}
}
