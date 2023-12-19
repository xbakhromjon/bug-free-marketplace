package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang-project-template/internal/basket/app"
	"net/http"
)

func NewRouter(cartService app.CartService) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cartController := NewCartController(cartService)

	r.Route("/cart", func(r chi.Router) {
		r.Post("/create", cartController.CreateCart)
	})

	return r
}
