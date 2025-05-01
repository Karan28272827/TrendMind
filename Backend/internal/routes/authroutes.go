package routes

import (
	"net/http"
	"server/handlers"

	"server/middleware"

	"github.com/go-chi/chi/v5"
)

func AuthRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.CORSMiddleware())
	r.Group(func(r chi.Router) { //testing JWT middleware and route protection
		r.Use(middleware.JWTMiddleware)
		r.Get("/profile", handlers.Profile)
	})

	r.Post("/register", handlers.Register)              //register route
	r.Post("/login", handlers.Login)                    //login route
	r.Post("/forgot-password", handlers.ForgotPassword) // request-password-reset
	r.Post("/reset-password", handlers.ResetPassword)

	r.Get("/google", handlers.GoogleLogin)
	r.Get("/google/callback", handlers.GoogleCallback)

	r.Get("/verify", handlers.VerifyEmail)

	r.Get("/github", handlers.GithubLogin)
	r.Get("/github/callback", handlers.GithubCallback)

	

	return r
}
