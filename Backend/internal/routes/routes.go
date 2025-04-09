package routes

import (
	"net/http"
	"server/handlers"

	"server/middleware"

	"github.com/go-chi/chi/v5"
)

func AuthRouter() http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTMiddleware)
		r.Get("/profile", handlers.Profile)
	})

	r.Post("/register", handlers.Register) //register route
	r.Post("/login", handlers.Login)       //login route
	return r
}
