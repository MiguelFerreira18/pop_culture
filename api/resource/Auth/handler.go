package auth

import (
	"encoding/json"
	"net/http"
	"pop_culture/api/resource/common/log"
	"pop_culture/util/ctx"
	"pop_culture/util/jwt"

	e "pop_culture/api/resource/common/err"

	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type AuthAPI struct {
	Logger     *zerolog.Logger
	Repository AuthRepository
	JWTAuth    *jwtauth.JWTAuth
}

func NewAuthAPI(logger *zerolog.Logger, db *gorm.DB, jwtAuth *jwtauth.JWTAuth) *AuthAPI {
	return &AuthAPI{
		Logger:     logger,
		Repository: *NewAuthRepository(db),
		JWTAuth:    jwtAuth,
	}
}

func (api AuthAPI) Login(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())

	form := &FormLogin{}
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		api.Logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	user, err := api.Repository.Login(form.Email, form.Password)
	if err != nil {
		api.Logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespLoginFailure)
		return
	}
	jwtToken := jwt.NewJWTToken(user.ID, *user.Email, user.Role.Name)
	_, jwtString, err := jwtToken.Encode(api.JWTAuth)
	if err != nil {
		api.Logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJWTTokenCreationFailure)
		return
	}

	w.Header().Add("Authorization", "Bearer"+*jwtString)
	w.WriteHeader(http.StatusOK)
	jwtLogin := NewLoginDTO("Login successful", *jwtString)
	if err := json.NewEncoder(w).Encode(&jwtLogin); err != nil {
		api.Logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}

}
