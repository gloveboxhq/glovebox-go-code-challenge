package main

import (
	"os"
	"strconv"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/joho/godotenv"
)

var browser *rod.Browser

func init() {
	// Loading .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error while loading .env file")
	}

	// Using os package to get the env variable
	showBrowser, err := strconv.ParseBool(os.Getenv("SHOW_BROWSER"))
	if err != nil {
		showBrowser = false
	}

	// Creating a new launcher
	launcher, err := launcher.New().Headless(!showBrowser).Launch()
	if err != nil {
		panic("Error while creating launcher")
	}

	// Launching a new browser
	browser = rod.New().ControlURL(launcher)
	if err := browser.Connect(); err != nil {
		panic("Error while launching browser")
	}
}

func main() {
	var p PolicyProvider = FacebookProvider{}
	p.Login(os.Getenv("FACEBOOK_USERNAME"), os.Getenv("FACEBOOK_PASSWORD"))
	defer browser.Close()
}
