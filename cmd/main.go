package main

import (
	"CarRentalService/pkg/app"
	
)

// @title Car Rental API
// @version 1.0
// @description API для системы аренды автомобилей
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := app.NewApp()
	app.Run()
	
}	