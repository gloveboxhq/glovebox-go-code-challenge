package providers

import (
	"testing"

	"github.com/cosmotek/carrierproxy-poc/config"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestPolicyProviders(rootTest testing.T) {
	type loginTest struct {
		InputUsername, InputPassword string
		ExpectedError                error
	}

	type policyProviderLoginTestSet struct {
		Provider PolicyProvider
		Tests    map[string]loginTest
	}

	// fetch config from env
	conf := config.Get()

	// since all providers implement the same interface, all providers
	// can share the same table-tests for their interface methods.
	providers := []policyProviderLoginTestSet{
		{
			// Provider: NetflixPolicyProvider{},
			Tests: map[string]loginTest{
				"happy path": loginTest{
					InputUsername: conf.NetflixUsername,
					InputPassword: conf.NetflixPassword,
					ExpectedError: nil,
				},
			},
		},
	}

	// range providers and run tests for each provider as a subtest
	for _, provider := range providers {

		// create subtest for provider (by name)
		rootTest.Run(provider.Provider.ProviderName(), func(providerSetTest *testing.T) {

			// range tests and exec as subtest (by test name)
			for testName, testParams := range provider.Tests {
				providerSetTest.Run(testName, func(test *testing.T) {

					// exec login method
					err := provider.Provider.Login(testParams.InputPassword, testParams.InputUsername)
					if cmp.Equal(err, testParams.ExpectedError, cmpopts.EquateErrors()) {
						test.Errorf("expected '%v' but got '%v'", testParams.ExpectedError, err)
						test.Fail()
					}
				})
			}
		})
	}
}
