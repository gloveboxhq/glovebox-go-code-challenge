package netflix

import (
	"io"

	"github.com/cosmotek/carrierproxy-poc/providers"
)

type NetflixPolicyProvider struct{}

func (n NetflixPolicyProvider) ProviderName() string {
	return "NETFLIX"
}

func (n NetflixPolicyProvider) Login(username, password string) error {
	return nil
}

func (n NetflixPolicyProvider) Policies() ([]providers.Policy, error) {
	return nil, nil
}

func (n NetflixPolicyProvider) DocumentDownload(downloadKey string) (io.ReadCloser, error) {
	return nil, nil
}
