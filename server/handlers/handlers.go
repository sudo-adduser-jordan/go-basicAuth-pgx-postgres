package handlers

import (
	"fmt"
	"log"
	"main/database"
	"main/utils"
	"net/http"
)

type User struct {
	Username string
	Password string
}

// Create Account
func Register(w http.ResponseWriter, r *http.Request) {
	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	hashed_password, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatal("Hash password error")
	}

	if database.ReadUser(user.Username) == user.Username {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, "User exists: "+user.Username)

	} else {
		if database.CreateUser(user.Username, hashed_password) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "User created: "+user.Username)
		}
	}
}

// Read Account - TODO: assign cookie
func Login(w http.ResponseWriter, r *http.Request) {
	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	hash := database.ReadPassword(user.Username)

	if utils.CheckPasswordHash(user.Password, hash) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "User logged in: "+user.Username)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Update Account
func Update(w http.ResponseWriter, r *http.Request) {
	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	new_user := &User{
		Username: r.FormValue("new_username"),
		Password: r.FormValue("new_passwrod"),
	}
	hashed_password, err := utils.HashPassword(new_user.Password)
	if err != nil {
		log.Fatal("Hash password error")
	}

	hash := database.ReadPassword(user.Username)
	if utils.CheckPasswordHash(user.Password, hash) {
		database.UpdateUser(user.Username, new_user.Username, hashed_password)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "User Updated: "+user.Username)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Delete Account
func Delete(w http.ResponseWriter, r *http.Request) {
	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	password := database.ReadPassword(user.Username)

	if utils.CheckPasswordHash(user.Password, password) {
		database.DeleteUser(user.Username)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
