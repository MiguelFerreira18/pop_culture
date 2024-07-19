package user

import (
	"encoding/json"
	"net/http"
	e "pop_culture/api/resource/common/err"
	"pop_culture/api/resource/common/log"
	"pop_culture/util/ctx"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type UserApi struct {
	logger     *zerolog.Logger
	repository *UserRepository
}

func NewUserApi(logger *zerolog.Logger, repository *gorm.DB) *UserApi {
	return &UserApi{
		logger:     logger,
		repository: NewRepository(repository),
	}
}

func (up *UserApi) Create(w http.ResponseWriter, r *http.Request) {

	reqID := ctx.RequestID(r.Context())
	form := &FormUser{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		up.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	user, err := NewUser(form.Name, form.Email, form.Password)
	if err != nil {
		up.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		return
	}
	user.ID = uuid.New()
	addUser, err := up.repository.Create(user)
	if err != nil {
		up.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}
	up.logger.Info().Str(log.KeyReqID, reqID).Str("id", addUser.ID.String()).Msg("New User was added")
	w.WriteHeader(http.StatusCreated)
}

func (up *UserApi) Read(w http.ResponseWriter, r *http.Request) {

	reqID := ctx.RequestID(r.Context())
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}
	user, err := up.repository.Read(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		up.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}
	dto := user.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		up.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}

}

func (up *UserApi) Update(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	form := &FormUser{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		up.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONDecodeFailure)
		return
	}

	user, err := NewUser(form.Name, form.Email, form.Password)
	if err != nil {
		up.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDomainRulesFailure)
		return
	}

	user.ID = id

	rows, err := up.repository.Update(user)
	if err != nil {
		up.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.RespDBDataUpdateFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	up.logger.Info().Str(log.KeyReqID, reqID).Str("id", id.String()).Msg("User updated")

}

func (up *UserApi) Delete(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	rows, err := up.repository.Delete(id)
	if err != nil {
		up.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataRemoveFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	up.logger.Info().Str(log.KeyReqID, reqID).Str("id", id.String()).Msg("USer deleted")

}
