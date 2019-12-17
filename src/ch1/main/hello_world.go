package main

import (
	"ch1/service"
	"fmt"
)

var appName = "accountservice"

func main() {
	// fmt.Print("Hello World!")
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767") // NEW
}
