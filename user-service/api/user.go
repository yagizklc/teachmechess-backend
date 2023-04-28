package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yagizklc/teachmechess-backend/user-service/pkg"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := pkg.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		user, err := pkg.GetUserByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	} else {
		username := r.URL.Query().Get("username")
		if username != "" {
			user, err := pkg.GetUserByUsername(username)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(user)
		} else {
			http.Error(w, "Provide an ID or a username", http.StatusBadRequest)
		}
	}
}
