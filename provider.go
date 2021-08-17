package main

type ProviderData struct {
	ID       string
	Location string
	Routes   []Route
}

type Route struct {
	Name string
	Path string
}
