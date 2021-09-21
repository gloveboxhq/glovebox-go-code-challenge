package netflix

import (
	"testing"

	"github.com/cosmotek/carrierproxy-poc/config"
	"github.com/cosmotek/carrierproxy-poc/providers"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestLogin(rootTest testing.T) {
	type loginTest struct {
		InputUsername, InputPassword string
		ExpectedError                error

		// optional provider instance. if nil, a new instance will be created
		Provider providers.PolicyProvider
	}

	// fetch config from env
	conf := config.Get()

	// since all providers implement the same interface, all providers
	// can share the same table-tests for their interface methods.
	tests := map[string]loginTest{
		"happy path": {
			InputUsername: conf.NetflixUsername,
			InputPassword: conf.NetflixPassword,
			ExpectedError: nil,
		},
	}

	// range tests and exec as subtest (by test name)
	for testName, testParams := range tests {
		rootTest.Run(testName, func(test *testing.T) {
			// if provider wasn't supplied (is nil), create new instance
			var provider providers.PolicyProvider = testParams.Provider
			if provider == nil {
				provider = NewProvider()
			}

			// exec login method
			err := provider.Login(testParams.InputPassword, testParams.InputUsername)
			if cmp.Equal(err, testParams.ExpectedError, cmpopts.EquateErrors()) {
				test.Errorf("expected '%v' but got '%v'", testParams.ExpectedError, err)
				test.Fail()
			}
		})
	}
}
