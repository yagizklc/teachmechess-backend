package pkg

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	users = []User{
		{ID: 1, Username: "user1", Password: "$2a$10$VUIfR5R5jKf5GZz5Dn20wOBtL4EBBbw4aeXcG7x.gC8eZj5ckgcOu"},
		{ID: 2, Username: "user2", Password: "$2a$10$VUIfR5R5jKf5GZz5Dn20wOBtL4EBBbw4aeXcG7x.gC8eZj5ckgcOu"},
	}
)

func GetAllUsers() ([]User, error) {
	return users, nil
}

func GetUserByID(id int) (*User, error) {
	for _, u := range users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("User not found")
}

func GetUserByUsername(username string) (*User, error) {
	for _, u := range users {
		if u.Username == username {
			return &u, nil
		}
	}
	return nil, errors.New("User not found")
}

func (u *User) VerifyPassword(password string) error {
	// Compare the provided password with the user's stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errors.New("incorrect password")
	}
	return nil
}
