package pkg

import (
	"errors"
	"os"
	"time"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/yagizklc/teachmechess-backend/user-service/pkg"
)

var (
	jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateAuthToken(userID int) (string, error) {
	// Generate a JWT token with the user ID as the payload
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Set token expiration time to 24 hours from now
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", errors.New("error generating auth token")
	}

	return tokenString, nil
}

func ValidateCredentials(username, password string) (*pkg.User, error) {
	user, err := getUserByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	err = user.VerifyPassword(password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// helper function to get user data from the user service
func getUserByUsername(username string) (*pkg.User, error) {
	userServiceURL := os.Getenv("USER_SERVICE_URL") // Get the User Service URL from the environment variable
	resp, err := http.Get(fmt.Sprintf("%s/users/username/%s", userServiceURL, username))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("user not found")
	}

	user := &pkg.User{}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
