package main

import (
	// "fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	// "encoding/json"
	"server/db"
	"server/internal/routes"
	"server/middleware"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	db.InitDB()

	r := chi.NewRouter()

	r.Use(middleware.CORSMiddleware())

	r.Mount("/auth", routes.AuthRouter())

	http.ListenAndServe(":8080", r)

}
