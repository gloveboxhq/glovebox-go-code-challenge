package controller

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	url               = `https://pkg.go.dev/time`
	clickBtnId        = `#example-After`
	retrievingElement = `textarea`
)

func TestInitScrape(t *testing.T) {

	result, err := InitScrape(url, clickBtnId, retrievingElement)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestUsingEnvVar(t *testing.T) {

	os.Setenv("ENV_VAR", "value")
	defer os.Unsetenv("ENV_VAR")

	actual := os.Getenv("ENV_VAR")
	assert.Equal(t, "value", actual)

}
