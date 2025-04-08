package routes

import (
	"net/http"
	"server/handlers"

	"github.com/go-chi/chi/v5"
)

func AuthRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)
	return r
}
