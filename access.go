package main

type AccessRequest struct {
	username string
	password string
}

type CarrierProxy struct {
	PolicyProvider
}

func generateProvider(provider string) ProviderData {
	url := getPolicyProviderSite(provider)
	// in the case of complex navigation it would be handy to have a list of routes for reference by the application
	// this example will not use that but I am providing an example of it any way
	routes := []Route{}
	providerMeta := ProviderData{provider, url, routes}
	return providerMeta
}
