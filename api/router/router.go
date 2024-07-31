package router

import (
	"net/http"
	attribute "pop_culture/api/resource/Attribute"
	auth "pop_culture/api/resource/Auth"
	media "pop_culture/api/resource/Media"
	mediatype "pop_culture/api/resource/MediaType"
	mediatypeattribute "pop_culture/api/resource/MediaTypeAttribute"
	role "pop_culture/api/resource/Role"
	user "pop_culture/api/resource/User"
	userinterests "pop_culture/api/resource/UserInterests"
	usermedia "pop_culture/api/resource/UserMedia"
	"pop_culture/api/resource/health"
	"pop_culture/api/router/middleware"
	requestlog "pop_culture/api/router/middleware/requestLog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func New(logger *zerolog.Logger, database *gorm.DB, jwtAuth *jwtauth.JWTAuth) *chi.Mux {
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

	// verifier := middleware.NewJWTVerifier(logger, database)

	r.Use(middleware.RequestID)
	r.Use(middleware.ContentTypeJSON)
	r.Get("/api", health.Read)
	r.Route("/v1", func(r chi.Router) {

		AuthApi := auth.NewAuthAPI(logger, database, jwtAuth)
		r.Method(http.MethodPost, "/login", requestlog.NewHandler(AuthApi.Login, logger))

		userAPI := user.NewUserApi(logger, database)
		r.Method(http.MethodPost, "/users", requestlog.NewHandler(userAPI.Create, logger))
		r.Method(http.MethodGet, "/users/{id}", requestlog.NewHandler(userAPI.Read, logger))
		r.Method(http.MethodPut, "/users/{id}", requestlog.NewHandler(userAPI.Update, logger))
		r.Method(http.MethodDelete, "/users/{id}", requestlog.NewHandler(userAPI.Delete, logger))

		mediaTypeApi := mediatype.NewMediaTypeAPI(logger, database)
		r.Method(http.MethodGet, "/mediatype", requestlog.NewHandler(mediaTypeApi.List, logger))
		r.Method(http.MethodPost, "/mediatype", requestlog.NewHandler(mediaTypeApi.Create, logger))
		r.Method(http.MethodGet, "/mediatype/{id}", requestlog.NewHandler(mediaTypeApi.Read, logger))
		r.Method(http.MethodPut, "/mediatype/{id}", requestlog.NewHandler(mediaTypeApi.Update, logger))
		r.Method(http.MethodDelete, "/mediatype/{id}", requestlog.NewHandler(mediaTypeApi.Delete, logger))

		mediaAPI := media.NewMediaAPI(logger, database)
		r.Method(http.MethodGet, "/media", requestlog.NewHandler(mediaAPI.List, logger))
		r.Method(http.MethodPost, "/media", requestlog.NewHandler(mediaAPI.Create, logger))
		r.Method(http.MethodGet, "/media/{id}", requestlog.NewHandler(mediaAPI.Read, logger))
		r.Method(http.MethodPut, "/media/{id}", requestlog.NewHandler(mediaAPI.Update, logger))
		r.Method(http.MethodDelete, "/media/{id}", requestlog.NewHandler(mediaAPI.Delete, logger))

		AttributeAPI := attribute.NewAttributeAPI(logger, database)
		r.Method(http.MethodGet, "/attribute", requestlog.NewHandler(AttributeAPI.List, logger))
		r.Method(http.MethodPost, "/attribute", requestlog.NewHandler(AttributeAPI.Create, logger))
		r.Method(http.MethodGet, "/attribute/{id}", requestlog.NewHandler(AttributeAPI.Read, logger))
		r.Method(http.MethodPut, "/attribute/{id}", requestlog.NewHandler(AttributeAPI.Update, logger))
		r.Method(http.MethodDelete, "/attribute/{id}", requestlog.NewHandler(AttributeAPI.Delete, logger))

		mediaTypeAttributeAPI := mediatypeattribute.NewMediaTypeAttributeAPI(logger, database)
		r.Method(http.MethodGet, "/mediatype/{id}/attribute", requestlog.NewHandler(mediaTypeAttributeAPI.GetInterestsFromUser, logger))
		r.Method(http.MethodPost, "/mediatype/{id}/attribute", requestlog.NewHandler(mediaTypeAttributeAPI.AppendAttribute, logger))
		r.Method(http.MethodDelete, "/mediatype/{id}/attribute", requestlog.NewHandler(mediaTypeAttributeAPI.RemoveAttribute, logger))

		userMediaAPI := usermedia.NewUserMediaAPI(logger, database)
		r.Method(http.MethodGet, "/user/{id}/media", requestlog.NewHandler(userMediaAPI.GetMediaFromUser, logger))
		r.Method(http.MethodPost, "/user/{id}/media", requestlog.NewHandler(userMediaAPI.AddMediaToUser, logger))
		r.Method(http.MethodDelete, "/user/{id}/media", requestlog.NewHandler(userMediaAPI.RemoveMediaFromUser, logger))

		userInterestAPI := userinterests.NewUserInterestAPI(logger, database)
		r.Method(http.MethodGet, "/user/{id}/interest", requestlog.NewHandler(userInterestAPI.GetInterestsFromUser, logger))
		r.Method(http.MethodPost, "/user/{id}/interest", requestlog.NewHandler(userInterestAPI.AppendInterestToUser, logger))
		r.Method(http.MethodDelete, "/user/{id}/interest", requestlog.NewHandler(userInterestAPI.RemoveInterestFromUser, logger))

		roleAPI := role.NewRoleAPI(logger, database)
		// r.With(middleware.RoleCheck(*verifier, "Admin", "User")).Route("/role", func(r chi.Router) {
		// 	r.Method(http.MethodPost, "/", requestlog.NewHandler(roleAPI.Create, logger))
		// 	r.Method(http.MethodGet, "/{id}", requestlog.NewHandler(roleAPI.Read, logger))
		// 	r.Method(http.MethodPut, "/{id}", requestlog.NewHandler(roleAPI.Update, logger))
		// 	r.Method(http.MethodDelete, "/{id}", requestlog.NewHandler(roleAPI.Delete, logger))
		//
		// })
		r.Method(http.MethodPost, "/role", requestlog.NewHandler(roleAPI.Create, logger))
		r.Method(http.MethodGet, "/role/{id}", requestlog.NewHandler(roleAPI.Read, logger))
		r.Method(http.MethodPut, "/role/{id}", requestlog.NewHandler(roleAPI.Update, logger))
		r.Method(http.MethodDelete, "/role/{id}", requestlog.NewHandler(roleAPI.Delete, logger))

	})

	return r
}
