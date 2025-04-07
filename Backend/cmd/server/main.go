package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"encoding/json"
	"server/db"
	"server/internal/routes"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	db.InitDB()

	r := chi.NewRouter()
	r.Mount("/auth", routes.AuthRouter())

	var users []User
	users = append(users, User{Name: "Soham", Age: 21})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		for i := range users {
			fmt.Fprintf(w, "Name: %s, Age: %d\n", users[i].Name, users[i].Age)
		}
	})

	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		fmt.Fprintf(w, "User id: %s\n", userId)
	})

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		users = append(users, User{Name: user.Name, Age: user.Age})
		fmt.Fprintf(w, "Successfully added user %s\n", user.Name)
	})

	http.ListenAndServe(":8080", r)

}
