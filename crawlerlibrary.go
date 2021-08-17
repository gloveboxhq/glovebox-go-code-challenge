package main

import "fmt"

func github(providerMeta ProviderData, requestMeta AccessRequest) {
	page := createPageConnection(providerMeta.Location, 10)
	a1 := fmt.Sprintf("images/%s-a-%s.png", providerMeta.ID, requestMeta.username)
	collectScreenshot(a1, page)
	clickElement("a.HeaderMenu-link.flex-shrink-0.no-underline", page)
	b1 := fmt.Sprintf("images/%s-b-%s.png", providerMeta.ID, requestMeta.username)
	collectScreenshot(b1, page)
	enterInput("input#login_field.form-control.input-block", requestMeta.username, page)
	enterInput("input#password.form-control.input-block", requestMeta.password, page)
	clickElement("input.btn.btn-primary.btn-block", page)
	c1 := fmt.Sprintf("images/%s-c-%s.png", providerMeta.ID, requestMeta.username)
	collectScreenshot(c1, page)
}

func humbleBundle(providerMeta ProviderData, requestMeta AccessRequest) {
	page := createPageConnection(providerMeta.Location, 10)
	a2 := fmt.Sprintf("images/%s-a-%s.png", providerMeta.ID, requestMeta.username)
	collectScreenshot(a2, page)
	clickElement("a.navbar-item.js-user-navbar-item.user-navbar-item.js-account-login.logged-out.desktop.button-title.navbar-login", page)
	b2 := fmt.Sprintf("images/%s-b-%s.png", providerMeta.ID, requestMeta.username)
	collectScreenshot(b2, page)
	enterInput("input.text-input", requestMeta.username, page)
	enterInput("input.text-input.js-input", requestMeta.password, page)
	clickElement("button.flat-cta-button.blue.no-style-button", page)
	c2 := fmt.Sprintf("images/%s-c-%s.png", providerMeta.ID, requestMeta.username)
	collectScreenshot(c2, page)
}
