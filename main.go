package main

import (
	"flag"

	"github.com/joho/godotenv"
)

func main() {
	providerID := flag.String("provider", "", "id of the policy providing entity")
	account := flag.String("account", "", "name of the users account")
	flag.Parse()
	godotenv.Load()
	var p PolicyProvider
	provider := generateProvider(*providerID)
	p = provider
	p.Login(*account, getAccountCredentials(*account))
}
