package netflix

import (
	"errors"
	"io"
	"log"

	"github.com/cosmotek/carrierproxy-poc/providers"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const (
	netflixAccountPageUrl = "https://netflix.com/YourAccount"
	noAccountFoundText    = "Sorry, we can't find an account with this email address. Please try again or"
	incorrectPasswordText = "Incorrect password."
)

var (
	ErrLoginIncorrectPassword = errors.New("failed to login, incorrect password provided")
	ErrLoginNoAccountFound    = errors.New("failed to login, no account found with provided email")
	ErrLoginUnknownReason     = errors.New("failed to login, reason unknown")
)

type NetflixPolicyProvider struct {
	browser *rod.Browser
}

func NewProvider() NetflixPolicyProvider {
	return NetflixPolicyProvider{
		browser: rod.New().MustConnect(),
	}
}

func (n NetflixPolicyProvider) ProviderName() string {
	return "NETFLIX"
}

func (n NetflixPolicyProvider) Login(username, password string) error {
	page, err := n.browser.Page(proto.TargetCreateTarget{URL: netflixAccountPageUrl})
	if err != nil {
		return err
	}

	err = page.WaitLoad()
	if err != nil {
		return err
	}

	info, err := page.Info()
	if err != nil {
		return err
	}

	log.Println(info.URL)

	if info.URL != "https://www.netflix.com/login?nextpage=https%3A%2F%2Fwww.netflix.com%2FYourAccount" {
		return errors.New("failed to load auth wall")
	}

	usernameInput, err := page.Element("userLoginId")
	if err != nil {
		return err
	}

	err = usernameInput.Input(username)
	if err != nil {
		return err
	}

	passwordInput, err := page.Element("password")
	if err != nil {
		return err
	}

	err = passwordInput.Input(password)
	if err != nil {
		return err
	}

	info, err = page.Info()
	if err != nil {
		return err
	}

	log.Println(info.URL)

	if info.URL != netflixAccountPageUrl {
		exists, _, err := page.HasR("div", noAccountFoundText)
		if err != nil {
			return err
		}

		if exists {
			return ErrLoginNoAccountFound
		}

		exists, _, err = page.HasR("b", incorrectPasswordText)
		if err != nil {
			return err
		}

		if exists {
			return ErrLoginIncorrectPassword
		}

		return ErrLoginUnknownReason
	}

	return nil
}

func (n NetflixPolicyProvider) Policies() ([]providers.Policy, error) {
	return nil, nil
}

func (n NetflixPolicyProvider) DocumentDownload(downloadKey string) (io.ReadCloser, error) {
	return nil, nil
}

func (n NetflixPolicyProvider) Close() error {
	return n.browser.Close()
}
