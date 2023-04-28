package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/yagizklc/teachmechess-backend/auth-service/api"
	"github.com/yagizklc/teachmechess-backend/auth-service/pkg"
)

func mockUserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestLogin(t *testing.T) {
	// Start the HTTP server
	server := httptest.NewServer(http.HandlerFunc(api.Login))
	defer server.Close()

	// Mock the user-service register endpoint
	mockRegisterServer := httptest.NewServer(http.HandlerFunc(mockUserRegisterHandler))
	defer mockRegisterServer.Close()

	// Create a new user account
	payload := []byte(`{"username":"testuser","password":"testpassword"}`)
	req, _ := http.NewRequest("POST", mockRegisterServer.URL+"/users/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %v, but got %v", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()

	// Attempt to log in with the same credentials
	payload = []byte(`{"username":"testuser","password":"testpassword"}`)
	req, _ = http.NewRequest("POST", server.URL+"/auth/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %v, but got %v", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()

	// Parse the token from the response body
	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)
	token := response["token"]

	// Load JwtSecretKey from the environment variable
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	// Verify that the token is valid by decoding it
	claims := &pkg.Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		t.Errorf("Error decoding token: %v", err)
	}
	if claims.UserID != 1 {
		t.Errorf("Expected user ID 1, but got %v", claims.UserID)
	}
}
