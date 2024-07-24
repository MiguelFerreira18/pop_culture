package userinterests

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

type UserInterestAPI struct {
	logger     *zerolog.Logger
	repository *UserInterestRepository
}

type formMediaTypeS struct {
	MediaTypeID uint `json:"id"`
}

func NewUserInterestAPI(logger *zerolog.Logger, db *gorm.DB) *UserInterestAPI {
	return &UserInterestAPI{
		logger:     logger,
		repository: NewUserInterestRepository(db),
	}
}

func (api UserInterestAPI) GetInterestsFromUser(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	medias, err := api.repository.GetInterestFromUser(userID)
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

func (api UserInterestAPI) AppendInterestToUser(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	form := formMediaTypeS{}
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONDecodeFailure)
		return
	}

	user, err := api.repository.AddInterestToUser(userID, form.MediaTypeID)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}

	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", user.ID.String()).Msg("Interest was added to the user")
	w.WriteHeader(http.StatusCreated)
}

func (api UserInterestAPI) RemoveInterestFromUser(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	userID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	form := formMediaTypeS{}
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONDecodeFailure)
		return
	}

	user, err := api.repository.RemoveInterestFromUser(userID, form.MediaTypeID)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataRemoveFailure)
		return
	}

	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", user.ID.String()).Msg("Interest removed from user")

}
