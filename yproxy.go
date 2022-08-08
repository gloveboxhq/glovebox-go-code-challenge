package main

import (
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

	/*
		proxy.loginUrl = "https://accounts.craigslist.org/login"
		proxy.usernameSelector = "#inputEmailHandle"
		proxy.passwordSelector = "#inputPassword"
		proxy.loginButtonSelector = "#login"
	*/
	proxy.loginUrl = "https://news.ycombinator.com/login?goto=newest"
	proxy.usernameSelector = "body > form:nth-child(4) > table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(2) > input:nth-child(1)"
	proxy.passwordSelector = "body > form:nth-child(4) > table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(2) > input:nth-child(1)"
	proxy.loginButtonSelector = "body > form:nth-child(4) > input[type=submit]:nth-child(4)"

	proxy.browser = rod.New().MustConnect().NoDefaultDevice()
	page := proxy.browser.MustPage(proxy.loginUrl).MustWindowFullscreen()

	page.MustElement(proxy.usernameSelector).MustInput(username)
	page.MustElement(proxy.passwordSelector).MustInput(password)
	page.MustElement(proxy.loginButtonSelector).MustClick()

	el := page.MustElement("#me")
	fmt.Println(el.MustText())
	return nil
}
