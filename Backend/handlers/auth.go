package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"server/db"
	"server/utils"

	"server/models"

	_ "github.com/lib/pq"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	var HashedPass = utils.HashPassword(user.Password)

	row := db.DB.QueryRow("SELECT email FROM users WHERE email=$1", user.Email)

	var duplicateEmail string
	duplicateErr := row.Scan(&duplicateEmail)
	fmt.Println(duplicateEmail)
	if duplicateErr == sql.ErrNoRows {
		fmt.Println("No common email")
	} else if duplicateErr != nil {
		fmt.Println("some unknown error", duplicateErr)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	} else {
		fmt.Println("email is duplicate Or ")
		fmt.Fprintln(w, "email is duplicate", user.Email)
		return
	}

	sqlQuery := `INSERT INTO users (name, email, password) VALUES($1, $2, $3)`
	var _, err = db.DB.Exec(sqlQuery, user.Name, user.Email, HashedPass)
	if err != nil {
		fmt.Println("Row not inserted ", err)
		http.Error(w, "Row was not inserted", http.StatusBadRequest)
		return
	} else {
		fmt.Println("\nRow inserted")
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser models.LoginUser

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var row = db.DB.QueryRow("SELECT password FROM users WHERE email=$1", loginUser.Email)
	var dbPassword string

	if err := row.Scan(&dbPassword); err != nil { //Inputting password from database into dbPassword variable
		fmt.Println("There was some error", err)
		return
	}

	fmt.Printf("HashedPass: %s and DBPass: %s", loginUser.Password, dbPassword)

	if !utils.CompareHashAndPassword(dbPassword, loginUser.Password) { //Checking if password matches
		fmt.Println("Invalid pass")
		fmt.Fprintln(w, "Invalid password")
		return
	} else {
		fmt.Println("\nCorrect password")
		fmt.Fprintln(w, "Correct password, logged in")
	}
}
