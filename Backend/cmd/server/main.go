package main

import (
	// "fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	// "encoding/json"
	"os"
	"server/db"
	"server/internal/routes"
	"server/middleware"
	"server/utils"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)


func main() {

	db.InitDB()
	utils.InitOAuth()
	gothic.Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

	r := chi.NewRouter()

	r.Use(middleware.CORSMiddleware())

	r.Mount("/auth", routes.AuthRouter())
	r.Mount("/chat", routes.ChatRouter())


	http.ListenAndServe(":8080", r)

}
