package main

import (
	"2022-08-08/glovebox-go-code-challenge/controller"
	"fmt"
)

func main() {
	//url = `https://pkg.go.dev/time`
	//clickBtnId = `#example-After`
	//retrievingElement = `textarea`
	var url, clickBtnId, retrievingElement string

	fmt.Print("Enter url: ")
	fmt.Scan(&url)

	fmt.Print("Enter clickBtnId: ")
	fmt.Scan(&clickBtnId)

	fmt.Print("Enter retrievingElement: ")
	fmt.Scan(&retrievingElement)

	controller.InitScrape(url, clickBtnId, retrievingElement)

}
