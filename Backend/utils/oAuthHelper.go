package utils

import (
	"database/sql"
	"fmt"
	"net/http"

	"server/db"
)

func HandleOAuthCallback(w http.ResponseWriter, name string, email string) {
	row := db.DB.QueryRow("SELECT email FROM users WHERE email=$1", email)
	existingEmail, err := IsDuplicateEmail(row)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if existingEmail == "" {
		// register user
		password := sql.NullString{String: "", Valid: false}
		_, err := db.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", name, email, password)
		if err != nil {
			http.Error(w, "User registration failed", http.StatusInternalServerError)
			return
		}
		fmt.Println("New user registered:", email)
	} else {
		fmt.Println("User already registered:", email)
	}

	// issue JWT
	token, err := CreateToken(name, email)
	if err != nil {
		http.Error(w, "JWT creation failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Logged in!\nJWT: %s", token)
}
