package netflix

import (
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/cosmotek/carrierproxy-poc/provider"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const (
	netflixAccountPageUrl   = "https://www.netflix.com/YourAccount"
	authWallUrlWithRedirect = "https://www.netflix.com/login?nextpage=https%3A%2F%2Fwww.netflix.com%2FYourAccount"

	noAccountFoundText    = "Sorry, we can't find an account with this email address. Please try again or"
	incorrectPasswordText = "Incorrect password."

	LOGGED_IN  LoginStatus = "LOGGED_IN"
	LOGGED_OUT LoginStatus = "LOGGED_OUT"
	UNKNOWN    LoginStatus = "UNKNOWN"
)

var (
	ErrLoginIncorrectPassword = errors.New("failed to login, incorrect password provided")
	ErrLoginNoAccountFound    = errors.New("failed to login, no account found with provided email")
	ErrLoginUnknownReason     = errors.New("failed to login, reason unknown")
	ErrFailureUnknownReason   = errors.New("failure due to unknown reason")
	ErrNotImplemented         = errors.New("method not implemented")
)

// LoginStatus is a type used to indicate login status within the context of this provider
type LoginStatus string

// NetflixPolicyProvider implements the provider.PolicyProvider interface
type NetflixPolicyProvider struct {
	browser *rod.Browser
}

// NewProvider creates a new provider instance by connecting
// to the automation browser and returning a provider with an
// active browser connection.
func NewProvider() (NetflixPolicyProvider, error) {
	browser := rod.New()
	err := browser.Connect()
	if err != nil {
		return NetflixPolicyProvider{}, err
	}

	return NetflixPolicyProvider{
		browser: browser,
	}, nil
}

// NewProviderFromBrowser creates and returns a new provider instance
// using the provided browser. This may be useful for tests where custom
// browser parameters may be required for debugging.
func NewProviderFromBrowser(browser *rod.Browser) NetflixPolicyProvider {
	return NetflixPolicyProvider{
		browser: browser,
	}
}

// GetLoginStatus attempts to determine the login status for the provider using the given
// username.
func (n NetflixPolicyProvider) GetLoginStatus(username string) (LoginStatus, error) {
	// load account page
	page, err := n.browser.Page(proto.TargetCreateTarget{URL: netflixAccountPageUrl})
	if err != nil {
		return UNKNOWN, err
	}
	defer page.Close()

	err = page.WaitLoad()
	if err != nil {
		return UNKNOWN, err
	}

	info, err := page.Info()
	if err != nil {
		return UNKNOWN, err
	}

	// check url to determine status
	switch info.URL {
	case authWallUrlWithRedirect:
		return LOGGED_OUT, nil
	case netflixAccountPageUrl:
		exists, elem, err := page.Has(".account-section-email")
		if err != nil {
			return UNKNOWN, err
		}

		text, err := elem.Text()
		if err != nil {
			return UNKNOWN, err
		}

		if !exists || text != username {
			return UNKNOWN, ErrFailureUnknownReason
		}

		return LOGGED_IN, nil
	default:
		return UNKNOWN, ErrFailureUnknownReason
	}
}

func (n NetflixPolicyProvider) Login(username, password string) error {
	// check login status
	status, err := n.GetLoginStatus(username)
	if err != nil {
		return err
	}

	// check if logged in already, if so, return nil
	if status == LOGGED_IN {
		return nil
	}

	// load account page expecting redirect to auth wall
	page, err := n.browser.Page(proto.TargetCreateTarget{URL: netflixAccountPageUrl})
	if err != nil {
		return err
	}
	defer page.Close()

	err = page.WaitLoad()
	if err != nil {
		return err
	}

	info, err := page.Info()
	if err != nil {
		return err
	}

	// verify auth wall was hit
	if info.URL != authWallUrlWithRedirect {
		log.Println("screenshoting page for later debugging")
		page.MustScreenshotFullPage(fmt.Sprintf("netflix_login_failure_%s.png", time.Now().String()))
		time.Sleep(time.Second * 3)

		return ErrLoginUnknownReason
	}

	// get username input element
	usernameInput, err := page.Element("[name=userLoginId]")
	if err != nil {
		return err
	}

	// input username
	err = usernameInput.Input(username)
	if err != nil {
		return err
	}

	// get password input element
	passwordInput, err := page.Element("[name=password]")
	if err != nil {
		return err
	}

	// input password
	err = passwordInput.Input(password)
	if err != nil {
		return err
	}

	// get submit button
	submitButton, err := page.Element("[type=submit]")
	if err != nil {
		return err
	}

	// press submit button
	err = submitButton.Click(proto.InputMouseButtonLeft)
	if err != nil {
		return err
	}

	err = page.WaitLoad()
	if err != nil {
		return err
	}

	// this blocks us just long enough to complete login,
	// otherwise we fail validation on data that is prematurely
	// fetched.
	log.Println("waiting for login to complete before validation")
	time.Sleep(time.Second * 5)

	info, err = page.Info()
	if err != nil {
		return err
	}

	if info.URL != netflixAccountPageUrl {
		exists, elem, err := page.HasR("div", noAccountFoundText)
		if err != nil {
			return err
		}

		visible, err := elem.Visible()
		if err != nil {
			return err
		}

		if exists && visible {
			return ErrLoginNoAccountFound
		}

		exists, elem, err = page.HasR("b", incorrectPasswordText)
		if err != nil {
			return err
		}

		visible, err = elem.Visible()
		if err != nil {
			return err
		}

		if exists && visible {
			return ErrLoginIncorrectPassword
		}

		log.Println("screenshoting page for later debugging")
		page.MustScreenshotFullPage(fmt.Sprintf("netflix_login_failure_%s.png", time.Now().String()))
		time.Sleep(time.Second * 3)

		return ErrLoginUnknownReason
	}

	exists, elem, err := page.Has(".account-section-email")
	if err != nil {
		return err
	}

	text, err := elem.Text()
	if err != nil {
		return err
	}

	// output account email for debugging (shown on account page)
	log.Println("account field contains:", text)

	if !exists || text != username {
		log.Println("screenshoting page for later debugging")
		page.MustScreenshotFullPage(fmt.Sprintf("netflix_login_failure_%s.png", time.Now().String()))
		time.Sleep(time.Second * 3)

		return ErrLoginUnknownReason
	}

	return nil
}

func (n NetflixPolicyProvider) Policies() ([]provider.Policy, error) {
	return nil, ErrNotImplemented
}

func (n NetflixPolicyProvider) DocumentDownload(downloadKey string) (io.ReadCloser, error) {
	return nil, ErrNotImplemented
}

// Close releases use of the browser instance
func (n NetflixPolicyProvider) Close() error {
	return n.browser.Close()
}
