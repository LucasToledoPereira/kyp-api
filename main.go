package main

import (
	"fmt"

	"github.com/LucasToledoPereira/kyp-api/src/cross_cutting/config"
	"github.com/LucasToledoPereira/kyp-api/src/cross_cutting/ioc"
)

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @BasePath /v1
// @title			Know Your Pet
// @version		1.0
// @description	Service to serve a pet information to the app.
// @BasePath		/
func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	fmt.Println("environment variables loaded successfully...")

	server, err := ioc.Bootstrap()

	if err != nil {
		panic(err)
	}
	fmt.Println("dependencies injected successfully")

	server.Run()
}
