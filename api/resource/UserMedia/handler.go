package usermedia

import (
	"encoding/json"
	"net/http"
	"pop_culture/util/ctx"

	e "pop_culture/api/resource/common/err"
	"pop_culture/api/resource/common/log"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type UserMediaAPI struct {
	logger     *zerolog.Logger
	repository *UserMediaRepository
}

type formMediaS struct {
	MediaID uint `json:"id"`
}

func NewUserMediaAPI(logger *zerolog.Logger, database *gorm.DB) *UserMediaAPI {
	return &UserMediaAPI{
		logger:     logger,
		repository: NewUserMediaRepository(database),
	}
}

func (api UserMediaAPI) GetMediaFromUser(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	medias, err := api.repository.GetMediasFromUser(userID)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	if err := json.NewEncoder(w).Encode(&medias); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}

}

func (api UserMediaAPI) AddMediaToUser(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}
	form := &formMediaS{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	user, err := api.repository.AddMediaToUser(form.MediaID, userID)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msgf("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}

	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", user.ID.String()).Msg("Media was added to user")
	w.WriteHeader(http.StatusCreated)

}

func (api UserMediaAPI) RemoveMediaFromUser(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}
	form := &formMediaS{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	user, err := api.repository.RemoveMediaFromUser(form.MediaID, userID)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msgf("")
		e.ServerError(w, e.RespDBDataRemoveFailure)
		return
	}

	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", user.ID.String()).Msg("Media was removed from user")

}
