package main

import (
	"log"
	"net/http"

	"github.com/yagizklc/teachmechess-backend/user-service/api"
)

func main() {
	http.HandleFunc("/users", api.GetUsers)
	http.HandleFunc("/users/:id", api.GetUserByID)
	log.Fatal(http.ListenAndServe(":8080", nil))
	// Replace ":8080" with the port number you want to use
}
