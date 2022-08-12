package Scrape

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func ScrapingFuncV1(url, clickBtnId, retrievingElement string) string {

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > footer`),
		// find and click "Example" link
		chromedp.Click(clickBtnId, chromedp.NodeVisible),
		// retrieve the text of the textarea
		chromedp.Value(clickBtnId+` `+retrievingElement, &example),
	)
	if err != nil {
		log.Fatal(err)
	}

	return example
}
