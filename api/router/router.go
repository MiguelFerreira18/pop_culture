package router

import (
	"pop_culture/api/resource/health"
	"pop_culture/api/router/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

const apiVersion = "/api/v1/"

func New(logger *zerolog.Logger, database *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api", health.Read)
	r.Route("/v1", func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.ContentTypeJSON)

	})
	return r
}
