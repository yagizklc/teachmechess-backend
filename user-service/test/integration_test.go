package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yagizklc/teachmechess-backend/user-service/api"
	"github.com/yagizklc/teachmechess-backend/user-service/pkg"
)

func TestGetUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(api.GetUsers))
	defer server.Close()

	resp, err := http.Get(server.URL + "/users")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
	}

	var users []pkg.User
	json.NewDecoder(resp.Body).Decode(&users)

	expectedUsers := 2
	if len(users) != expectedUsers {
		t.Errorf("Expected %v users, got %v", expectedUsers, len(users))
	}
}

func TestGetUserByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(api.GetUserByID))
	defer server.Close()

	// Test GetUserByID
	resp, err := http.Get(server.URL + "/users?id=1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
	}

	var user pkg.User
	json.NewDecoder(resp.Body).Decode(&user)

	expectedID := 1
	if user.ID != expectedID {
		t.Errorf("Expected user ID %v, got %v", expectedID, user.ID)
	}
}

func TestGetUserByUsername(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(api.GetUserByID))
	defer server.Close()

	// Test GetUserByUsername
	resp, err := http.Get(server.URL + "/users?username=user1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
	}

	var user pkg.User
	json.NewDecoder(resp.Body).Decode(&user)

	expectedUsername := "user1"
	if user.Username != expectedUsername {
		t.Errorf("Expected username %v, got %v", expectedUsername, user.Username)
	}
}
