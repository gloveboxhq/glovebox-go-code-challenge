package main

import (
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
)

const (
	host   = "www.facebook.com"
	scheme = "https"
)

var (
	baseUrl = &url.URL{
		Host:   host,
		Scheme: scheme,
	}
)

type FacebookProvider struct{}

type LoginResponse struct {
	IsLogin bool   `json:"isLogin"`
	Message string `json:"message"`
}

func (fb FacebookProvider) Login(username, password string) error {
	// Creating a new page
	page, err := browser.Page(proto.TargetCreateTarget{URL: baseUrl.String()})
	if err != nil {
		fmt.Println("Error while creating page : %w", err)
	}

	// Waiting for page load
	err = page.
		Timeout(time.Second * 30).
		WaitLoad()
	if err != nil {
		fmt.Println("Error while loading page: %w", err)
	}

	// Finding username element
	usernameEl, err := page.
		Timeout(time.Second * 15).
		Element(`input[id="email"]`)
	if err != nil {
		fmt.Println("Error finding username: %w", err)
	}

	// Inserting username in username textbox
	err = usernameEl.
		Timeout(time.Second * 15).
		Input(username)
	if err != nil {
		fmt.Println("Error inserting username: %w", err)
	}

	// Finding password element
	passwordEl, err := page.
		Timeout(time.Second * 15).
		Element(`input[id="pass"]`)
	if err != nil {
		fmt.Println("Error finding password: %w", err)
	}

	// Inserting password in password textbox
	err = passwordEl.
		Timeout(time.Second * 15).
		Input(password)
	if err != nil {
		fmt.Println("Error inserting password: %w", err)
	}

	// Finding login button
	loginEl, err := page.
		Timeout(time.Second * 15).
		Element(`button[type="submit"]`)
	if err != nil {
		fmt.Println("Error finding login button: %w", err)
	}

	// Pressing login button
	err = loginEl.
		Timeout(time.Second * 15).
		Type(input.Enter)
	if err != nil {
		fmt.Println("Error pressing login button: %w", err)
	}

	// Waiting for page load
	err = page.Timeout(time.Second * 30).WaitLoad()
	if err != nil {
		fmt.Println("Error while loading page: %w", err)
	}

	// Creating loginSuccessResponse channel to hold success response from CheckLoginSuccess function
	loginSuccessResponse := make(chan LoginResponse)

	// Creating loginFailureResponse channel to hold failure response from CheckLoginFailure function
	loginFailureResponse := make(chan LoginResponse)

	// Functions call with goroutines for success and failure
	go CheckLoginSuccess(page, loginSuccessResponse)
	go CheckLoginFailure(page, loginFailureResponse)

	// Selecting and executing channel
	select {
	case <-time.After(time.Second * 45):
		return fmt.Errorf("Login failed: Timeout")
	case successResponse := <-loginSuccessResponse:
		if successResponse.IsLogin {
			fmt.Println(successResponse.Message)
			return nil
		}
	case failedResponse := <-loginFailureResponse:
		if !failedResponse.IsLogin {
			return fmt.Errorf("Login failed. Error: %v", failedResponse.Message)

		}
	}
	return nil
}

// Goroutine to send loginSuccessResponse to channel
func CheckLoginSuccess(page *rod.Page, loginSuccessResponse chan LoginResponse) {
	loginSuccessEl, err := page.
		Timeout(time.Second * 45).
		Element(`span[dir="auto"]`)
	if err != nil {
		fmt.Println("Error finding logout button: %w", err)
	}

	if loginSuccessEl != nil {
		loginSuccessResponse <- LoginResponse{true, "Successful Login"}
	}

}

// Goroutine to send loginFailureResponse to channel
func CheckLoginFailure(page *rod.Page, loginFailureResponse chan LoginResponse) {
	_, err := page.
		Timeout(time.Second*10).
		ElementR("div", "you entered isn’t connected to an account.")

	if err == nil {
		loginFailureResponse <- LoginResponse{
			IsLogin: false,
			Message: "The email or mobile number is incorrect",
		}
	}

	_, err = page.
		Timeout(time.Second*10).
		ElementR("div", "you’ve entered is incorrect.")

	if err == nil {
		loginFailureResponse <- LoginResponse{
			IsLogin: false,
			Message: "The password is incorrect",
		}
	}

}

func (fb FacebookProvider) Policies() ([]Policy, error) {
	return nil, nil
}

func (fb FacebookProvider) DocumentDownload(downloadKey string) (io.ReadCloser, error) {
	return nil, nil
}
