package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"server/db"
	"server/utils"

	"server/models"

	_ "github.com/lib/pq"
	"github.com/markbates/goth/gothic"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	var HashedPass = utils.HashPassword(user.Password)

	row := db.DB.QueryRow("SELECT email FROM users WHERE email=$1", user.Email)

	email, duplicateErr := utils.IsDuplicateEmail(row)
	if duplicateErr != nil {
		fmt.Println("some unknown error", duplicateErr)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	} else if email != "" {
		fmt.Println("email is duplicate Or ")
		fmt.Fprintln(w, "email is duplicate", email)
		return
	} else {
		fmt.Println("No common email")
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

	var row = db.DB.QueryRow("SELECT name, password FROM users WHERE email=$1", loginUser.Email)
	var dbPassword string
	var dbName string

	if err := row.Scan(&dbName, &dbPassword); err != nil { //Inputting password from database into dbPassword variable
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
		tokenStr, err := utils.CreateToken(dbName, loginUser.Email)
		if err != nil {
			fmt.Println("There was some JWT token error", err)
			return
		}
		fmt.Fprintf(w, "Correct password, logged in\n JWT: %s", tokenStr)
	}
}

func Profile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test")
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) { //Function that initiates the redirection to google login
	r.URL.RawQuery = r.URL.RawQuery + "&provider=google"
	gothic.BeginAuthHandler(w, r)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		redirectUrl := fmt.Sprintf("%s/login?msg=LoginFailed", os.Getenv("FRONTEND_URL"))
		http.Redirect(w, r, redirectUrl, http.StatusUnauthorized)
		return
	}

	utils.HandleOAuthCallback(w, user.Name, user.Email, r)
}

func GithubLogin(w http.ResponseWriter, r *http.Request) {
	r.URL.RawQuery = r.URL.RawQuery + "&provider=github"
	gothic.BeginAuthHandler(w, r)
}

func GithubCallback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	utils.HandleOAuthCallback(w, user.Name, user.Email, r)
}
