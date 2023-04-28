package main

import (
	"log"
	"net/http"

	"github.com/yagizklc/teachmechess-backend/auth-service/api"
)

func main() {
	http.HandleFunc("/auth", api.Login)
	log.Fatal(http.ListenAndServe(":8080", nil))
	// Replace ":8080" with the port number you want to use
}
