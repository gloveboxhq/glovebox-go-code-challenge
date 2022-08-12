package controller

import (
	"2022-08-08/glovebox-go-code-challenge/models"
	"fmt"
	"io"
)

type Policy models.Policy

type PolicyProvider interface {
	Login(username, password string) error
	Policies() ([]Policy, error)
	DocumentDownload(downloadKey string) (io.ReadCloser, error)
}

func (p *Policy) Login(username, password string) error {
	fmt.Println("Login : ", username, password, p.CarrierID, p.PolicyNumber)
	return nil
}

func (p *Policy) DocumentDownload(downloadKey string) (io.ReadCloser, error) {
	fmt.Println("DocumentDownload : ", p.CarrierID, p.PolicyNumber, downloadKey)
	return nil, nil
}

func (p *Policy) Policies() ([]Policy, error) {
	fmt.Println("Policies : ", p.PolicyNumber, p.CarrierID)
	return nil, nil
}

func initiateLogin(p PolicyProvider, username, password string) error {
	return p.Login(username, password)
}

func initiatePolicies(p PolicyProvider) ([]Policy, error) {
	return p.Policies()
}

func initiateDocumentDownload(p PolicyProvider, downloadKey string) (io.ReadCloser, error) {
	return p.DocumentDownload(downloadKey)
}

func Init() {
	policy := &Policy{
		CarrierID:    "1234",
		PolicyNumber: "5678",
	}

	var userName, password string

	fmt.Print("Enter UserName: ")
	fmt.Scan(&userName)

	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	err := initiateLogin(policy, userName, password)
	if err != nil {
		return
	}

	download, err2 := initiateDocumentDownload(policy, "dl_key1234")
	fmt.Println("download resp : ", download)
	if err2 != nil {
		return
	}

	policies, err3 := initiatePolicies(policy)
	fmt.Println("policies resp : ", policies)
	if err3 != nil {
		return
	}

}
