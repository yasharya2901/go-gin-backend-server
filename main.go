package main

import (
	"awesomeProject/config"
	"awesomeProject/routes"
)

func main() {
	config.ConnectDatabase()

	router := routes.SetupRouter()

	router.Run(":8080")
}
