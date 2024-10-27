package main

import (
	"net/http"

	"github.com/Afomiat/GoCrudChallange/Database"
	"github.com/Afomiat/GoCrudChallange/Delivery/routers"
	"github.com/gorilla/handlers"
)
func main() {
	Database.DBconnection()
	router := routers.SetupRouter()

	allowedOrigins := []string{"http://localhost:5173"}

    corsHandler := handlers.CORS(
        handlers.AllowedOrigins(allowedOrigins),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )(router)

    http.ListenAndServe(":8080", corsHandler)
}