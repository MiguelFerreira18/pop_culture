package middleware

import (
	"net/http"
	auth "pop_culture/api/resource/Auth"
	"pop_culture/api/resource/common/log"
	"pop_culture/util/ctx"
	jwtutil "pop_culture/util/jwt"

	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type contextKey struct {
	name string
}
type JwtVerifier struct {
	Logger     *zerolog.Logger
	Repository auth.AuthRepository
}

func NewJWTVerifier(logger *zerolog.Logger, db *gorm.DB) *JwtVerifier {
	return &JwtVerifier{
		Logger:     logger,
		Repository: *auth.NewAuthRepository(db),
	}

}

var (
	ctxKey = &contextKey{"Token"}
)

func RoleCheck(verifier JwtVerifier, requiredRoles ...string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := ctx.RequestID(r.Context())

			token, _, err := jwtauth.FromContext(r.Context())
			if err != nil {
				verifier.Logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			userID, email, role, err := jwtutil.Decode(token)
			if err != nil {
				verifier.Logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			userBody, err := verifier.Repository.LoadUser(*userID, *email)
			if err != nil {
				verifier.Logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			for _, requiredRole := range requiredRoles {
				if requiredRole == *role && userBody.Role.Name == requiredRole {
					verifier.Logger.Info().Str(log.KeyReqID, reqID).Str(userBody.ID.String(), *role).Msg("User is correctly authenticated")
					h.ServeHTTP(w, r)
					return
				}
			}
			verifier.Logger.Error().Str(log.KeyReqID, reqID).Msg("Not Authorized")
			http.Error(w, http.StatusText(403), http.StatusUnauthorized)

		})
	}
}
