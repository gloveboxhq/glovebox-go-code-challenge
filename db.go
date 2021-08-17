package main

import (
	"os"
)

func getPolicyProviderSite(policyProvider string) string {
	// this would be where you would make a db connection to retrieve the latest and greatest policy holder data
	// since we aren't setting up a full db and just pulling env variables, we have this serving as a placeholder for that
	policySite := os.Getenv(policyProvider)
	return policySite
}

func getAccountCredentials(account string) string {
	// this would be where you would put a world class authentication and obfuscation system in place to avoid leaking users credentials
	// while preferming the same db look up that would be used to perform policy holder metadata retrieval
	accountCreds := os.Getenv(account)
	return accountCreds
}
