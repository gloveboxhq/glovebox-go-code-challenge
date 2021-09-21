package netflix

import (
	"testing"

	"github.com/cosmotek/carrierproxy-poc/config"
	"github.com/cosmotek/carrierproxy-poc/provider"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func createVisibleBrowser() *rod.Browser {
	browserStartOpts := launcher.New().
		Headless(false).
		Devtools(false)

	return rod.New().
		ControlURL(browserStartOpts.MustLaunch()).
		Trace(true).
		MustConnect()
}

func TestLogin(rootTest *testing.T) {
	type loginTest struct {
		InputUsername, InputPassword string
		ExpectedError                error

		// optional provider instance. if nil, a new instance will be created
		Provider provider.PolicyProvider
	}

	// fetch config from env
	conf := config.Get()

	// create a simple table for tests with test name
	// as map key and params as map value
	tests := map[string]loginTest{
		"happy path": {
			InputUsername: conf.NetflixUsername,
			InputPassword: conf.NetflixPassword,
			ExpectedError: nil,
		},
		"bad username": {
			InputUsername: conf.NetflixUsername[:len(conf.NetflixUsername)-1], // trim one char
			InputPassword: conf.NetflixPassword,
			ExpectedError: ErrLoginNoAccountFound,
		},
		"bad password": {
			InputUsername: conf.NetflixUsername,
			InputPassword: conf.NetflixPassword[:len(conf.NetflixPassword)-1], // trim one char
			ExpectedError: ErrLoginIncorrectPassword,
		},
	}

	// range tests and exec as subtest (by test name)
	for testName, testParams := range tests {
		rootTest.Run(testName, func(test *testing.T) {
			// if provider wasn't supplied (is nil), create new instance
			var provider provider.PolicyProvider = testParams.Provider
			if provider == nil {
				provider = NewProviderFromBrowser(createVisibleBrowser())
				defer provider.Close()
			}

			// exec login method
			err := provider.Login(testParams.InputUsername, testParams.InputPassword)
			if !cmp.Equal(err, testParams.ExpectedError, cmpopts.EquateErrors()) {
				test.Errorf("expected '%v' but got '%v'", testParams.ExpectedError, err)
				test.Fail()
			}
		})
	}
}
