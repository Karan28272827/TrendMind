package routes

import (
	"net/http"
	"server/handlers"
	"server/db"
	"server/middleware"

	"github.com/go-chi/chi/v5"
)

func ChatRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.CORSMiddleware())

	r.Post("/store-query", handlers.StoreUserQueryHandler(db.DB))
	r.Get("/products", handlers.GetProductsHandler(db.DB))

	return r
}