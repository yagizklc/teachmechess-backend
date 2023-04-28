package api

import (
	"encoding/json"
	"net/http"

	"github.com/yagizklc/teachmechess-backend/auth-service/pkg"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a LoginRequest struct
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the user's credentials
	user, err := pkg.ValidateCredentials(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate a new authentication token for the user
	token, err := pkg.GenerateAuthToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the authentication token to the client
	res := LoginResponse{Token: token}
	json.NewEncoder(w).Encode(res)
}
