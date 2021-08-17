package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod"
)

func crawler(providerMeta ProviderData, requestMeta AccessRequest) {
	// here we reference the crawler map to determine the correct crawler to run
	// I have two preconfigured crawlers set up for github and humblebundle
	// you can switch between the two by changing which provider you specify in the cli
	ProviderMap[providerMeta.ID].(func(ProviderData, AccessRequest))(providerMeta, requestMeta)
}

func errorHandler(report error) {
	// in a more robust setup I would want a custom global error handler to control application failure states
	// here we just print the error out because global error handlers are boring to look at
	fmt.Println(report)
	// force killing the app for good measure, a more complete solution would involve cleanup and more thorough
	// exit code handling for integration with extrenal monitoring tools
	os.Exit(1)
}

func createPageConnection(location string, timeout time.Duration) rod.Page {
	// this page connection features a timeout functionality that can be leveraged to clear out
	// long running instances and alert support staff of a potential outage on the provider side
	page := rod.New().MustConnect().MustPage()
	err := rod.Try(func() {
		page.Timeout(timeout * time.Second).MustNavigate(location)
	})
	if errors.Is(err, context.DeadlineExceeded) {
		// making a log entry and closing the application without throwing an error
		// this would could be used to call a custom alert setup for system outages
		fmt.Println("Scrapper failed to connect.")
		os.Exit(1)
	} else if err != nil {
		errorHandler(err)
	}
	return *page
}

func collectScreenshot(fileName string, page rod.Page) {
	page.MustWaitLoad().MustScreenshot(fileName)
}

func clickElement(element string, page rod.Page) {
	// here is the first selection where we are abstracting away some complexity from
	// from the "script" that will hold the core scrapper logic
	// the error handling here could be made arbitrarily complex without affecting
	// future devs updating the core workflow
	err := rod.Try(func() {
		page.MustElement(element).MustClick()
	})
	if err != nil {
		errorHandler(err)
	}
}

func enterInput(element string, input string, page rod.Page) {
	// it would make more sense to me that input would be the first arg
	// but in many of my current projects we promote collecting args in order of use
	// keeping that convention here is just an arbitrary choice I am making
	err := rod.Try(func() {
		page.MustElement(element).MustInput(input)
	})
	if err != nil {
		errorHandler(err)
	}
}
