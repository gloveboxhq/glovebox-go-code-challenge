package main

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
)

type PolicyProvider interface {
	Login(username, password string) error
	Policies() ([]Policy, error)
	DocumentDownload(downloadKey string) (io.ReadCloser, error)
}

func (providerMeta ProviderData) Login(username string, password string) error {
	var requestMeta AccessRequest
	requestMeta.username = username
	requestMeta.password = password
	// basic password length check, stricter controls on this processes
	// would be warranted if there were concerns about long running orphaned processes
	if len(requestMeta.password) < 1 {
		return errors.New("missing password")
	}
	ProviderMap[providerMeta.ID].(func(ProviderData, AccessRequest))(providerMeta, requestMeta)
	return nil
}

func (providerMeta ProviderData) Policies() ([]Policy, error) {
	// a stub to mock funcationality not required for the prototype
	return []Policy{}, nil
}

func (providerMeta ProviderData) DocumentDownload(downloadKey string) (io.ReadCloser, error) {
	// a stub to mock functionality not required for the prototype
	r := ioutil.NopCloser(strings.NewReader(downloadKey))
	return r, nil
}
