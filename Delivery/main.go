package main

import (
	"github.com/Afomiat/GoCrudChallange/Delivery/routers"
	"github.com/Afomiat/GoCrudChallange/Database"
)
func main() {
	Database.DBconnection()
	router := routers.SetupRouter()

	router.Run(":8080")
}