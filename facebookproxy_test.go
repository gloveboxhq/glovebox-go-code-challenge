package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	var p PolicyProvider = FacebookProvider{}
	err := p.Login(os.Getenv("FACEBOOK_USERNAME"), os.Getenv("FACEBOOK_PASSWORD"))
	assert.NoError(t, err)
}
