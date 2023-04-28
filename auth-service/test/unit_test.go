// auth_test.go
package main

import (
	"testing"

	"github.com/yagizklc/teachmechess-backend/auth-service/pkg"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestValidateCredentials(t *testing.T) {
	// Set up a mock user object
	user := User{
		Username: "testuser",
		Password: "$2a$10$OJXH8W1jsq5G5f5zlBZ0xuAQ0R7pgYvJ",
	}

	// Test a valid login
	validUser, err := pkg.ValidateCredentials("testuser", "testpassword")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if validUser.Username != user.Username {
		t.Errorf("Expected username %v, but got %v", user.Username, validUser.Username)
	}

	// Test an invalid login
	_, err = pkg.ValidateCredentials("testuser", "wrongpassword")
	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}
}
