package controller

import (
	"2022-08-08/glovebox-go-code-challenge/models"
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

type Scrape models.Scrape

type ScrapeProvider interface {
	ScrapingFuncV1() (string, error)
}

func (s *Scrape) ScrapingFuncV1() (string, error) {
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
		chromedp.Navigate(s.URL),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > footer`),
		// find and click "Example" link
		chromedp.Click(s.ClickBtnId, chromedp.NodeVisible),
		// retrieve the text of the textarea
		chromedp.Value(s.ClickBtnId+` `+s.RetrievingElement, &example),
	)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return example, nil
}

func InitiateScrapingFuncV1(p ScrapeProvider) (string, error) {
	return p.ScrapingFuncV1()
}

func InitScrape(URL, ClickBtnId, RetrievingElement string) (string, error) {
	scrape := &Scrape{
		URL:               URL,
		ClickBtnId:        ClickBtnId,
		RetrievingElement: RetrievingElement,
	}

	resp, err := InitiateScrapingFuncV1(scrape)
	if err != nil {
		return "", err
	}
	log.Printf("Extracted Data:\n%s", resp)

	return resp, nil
}
