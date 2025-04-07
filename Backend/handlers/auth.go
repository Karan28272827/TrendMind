package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/db"
	"server/utils"

	"server/models"

	_ "github.com/lib/pq"
)

// type HandlerUser struct {
// 	name     string `json:"name"`
// 	email    string `json:"email"`
// 	password string `json:"password"`
// }

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var HashedPass = utils.HashPassword(user.Password)

	sqlQuery := `INSERT INTO users (name, email, password) VALUES($1, $2, $3)`
	var _, err = db.DB.Exec(sqlQuery, user.Name, user.Email, HashedPass)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nRow inserted")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser models.LoginUser

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var hashedPass = utils.HashPassword(loginUser.Password)

	// sqlQuery := ``
}
