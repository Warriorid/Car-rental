// @title Car Rental Service API
// @version 1.0
// @description API for car rental management system
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

package main

import (
	"CarRentalService/pkg/app"
	
)

func main() {
	app := app.NewApp()
	app.Run()
	
}	