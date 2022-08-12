# carrierproxy-poc

This repository is used for creating a proof of concept around the `PolicyProvider` interface.

## Task

- [ ] Fork this repository and implement the `PolicyProvider`'s `Login` method against a web property of your choice using one of the scraping libraries listed below. 
- [ ] Implement tests that accept environment variables for the credentials (login & password). 
- [ ] Document your code and usage instructions.
- [ ] Send in a pull request for code review.

## Scraping Libraries

* [go-rod](https://github.com/go-rod/rod)
* [chromedp](https://github.com/chromedp/chromedp)

## Instructions to run the project.

### Run the project with:
#### This will execute a function that will scrape from `https://pkg.go.dev/time` after clicking the button `#example-After`
```
go run main.go
```

#### Example URL: `https://pkg.go.dev/time`
#### Example clickBtnId: `#example-After`
#### Example retrievingElement: `textarea`

## Run test cases. First move to the relevant directory.
```
cd controller
```


### There are two test cases:
- Test Scraping function

- Test Environment variables
```
go test -v
```
