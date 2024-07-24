package mediatypeattribute

import (
	"encoding/json"
	"net/http"
	"pop_culture/api/resource/common/log"
	"pop_culture/util/ctx"
	"strconv"

	e "pop_culture/api/resource/common/err"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type MediaTypeAttributeAPI struct {
	logger     *zerolog.Logger
	repository *MediaTypeAttributeRepository
}
type formAttributeS struct {
	AttributeID uint `json:"id"`
}

func NewMediaTypeAttributeAPI(logger *zerolog.Logger, database *gorm.DB) *MediaTypeAttributeAPI {
	return &MediaTypeAttributeAPI{
		logger:     logger,
		repository: NewMediaTypeAttributeRepository(database),
	}
}
func (api MediaTypeAttributeAPI) GetInterestsFromUser(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	mediaTypeID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	medias, err := api.repository.GetAttributesFromMediaType(uint(mediaTypeID))
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

func (api MediaTypeAttributeAPI) AppendAttribute(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	mediaTypeID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}
	form := &formAttributeS{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	mediaType, err := api.repository.AddAttribute(uint(mediaTypeID), uint(form.AttributeID))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msgf("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}
	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", strconv.FormatUint(uint64(mediaType.ID), 10)).Msg("Attribute was added to media type")
	w.WriteHeader(http.StatusCreated)
}

func (api MediaTypeAttributeAPI) RemoveAttribute(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	mediaTypeID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}
	form := &formAttributeS{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	mediaType, err := api.repository.RemoveAttribute(uint(mediaTypeID), uint(form.AttributeID))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msgf("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}
	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", strconv.FormatUint(uint64(mediaType.ID), 10)).Msg("Attribute was removed from media type")
}
