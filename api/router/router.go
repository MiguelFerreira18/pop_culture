package router

import (
	"net/http"
	mediatype "pop_culture/api/resource/MediaType"
	user "pop_culture/api/resource/User"
	"pop_culture/api/resource/health"
	"pop_culture/api/router/middleware"
	requestlog "pop_culture/api/router/middleware/requestLog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

const apiVersion = "/api/v1/"

func New(logger *zerolog.Logger, database *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.ContentTypeJSON)
	r.Get("/api", health.Read)
	r.Route("/v1", func(r chi.Router) {
		userAPI := user.NewUserApi(logger, database)
		r.Method(http.MethodPost, "/users", requestlog.NewHandler(userAPI.Create, logger))
		r.Method(http.MethodGet, "/users/{id}", requestlog.NewHandler(userAPI.Read, logger))
		r.Method(http.MethodPut, "/users/{id}", requestlog.NewHandler(userAPI.Update, logger))
		r.Method(http.MethodDelete, "/users/{id}", requestlog.NewHandler(userAPI.Delete, logger))

		mediaTypeApi := mediatype.NewMediaTypeAPI(logger, database)
		r.Method(http.MethodPost, "/mediatype", requestlog.NewHandler(mediaTypeApi.Create, logger))
		r.Method(http.MethodGet, "/mediatype/{id}", requestlog.NewHandler(mediaTypeApi.Read, logger))
		r.Method(http.MethodPut, "/mediatype/{id}", requestlog.NewHandler(mediaTypeApi.Update, logger))
		r.Method(http.MethodDelete, "/mediatype/{id}", requestlog.NewHandler(mediaTypeApi.Delete, logger))

	})

	return r
}
