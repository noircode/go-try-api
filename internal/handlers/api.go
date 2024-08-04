package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/noircode/go-try-api/internal/middleware"
)

func Handler(r *chi.Mux) {
	// global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(r chi.Router) {
		r.Use(middleware.Authorization)

		r.Get("/coins", GetCoinBalance)
	})
}
