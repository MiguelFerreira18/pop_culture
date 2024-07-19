package mediatype

import (
	"encoding/json"
	"net/http"
	e "pop_culture/api/resource/common/err"
	"pop_culture/api/resource/common/log"
	"pop_culture/util/ctx"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type MediaTypeAPI struct {
	logger     *zerolog.Logger
	repository *MediaTypeRepository
}

func NewMediaTypeAPI(logger *zerolog.Logger, database *gorm.DB) *MediaTypeAPI {
	return &MediaTypeAPI{
		logger:     logger,
		repository: NewMediaTypeRepository(database),
	}
}

func (api MediaTypeAPI) Create(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	form := &TypeMediaForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}
	mediaType, err := NewTypeMedia(form.Name)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msgf("")
		return
	}
	addMediaType, err := api.repository.Create(mediaType)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msgf("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}
	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", strconv.FormatUint(uint64(addMediaType.ID), 10)).Msg("New Media Type was added")
	w.WriteHeader(http.StatusCreated)
}

func (api MediaTypeAPI) Read(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	intId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}
	media, err := api.repository.Read(uint(intId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	dto := media.ToDTO()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}

}

func (api MediaTypeAPI) Update(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	intId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}
	form := &TypeMediaForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONDecodeFailure)
		return
	}

	media, err := NewTypeMedia(form.Name)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDomainRulesFailure)
		return
	}
	media.ID = uint(intId)

	rows, err := api.repository.Update(media)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataUpdateFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", strconv.FormatUint(uint64(intId), 10)).Msg("Media type updated")
}

func (api MediaTypeAPI) Delete(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())

	intId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	rows, err := api.repository.Delete(uint(intId))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataRemoveFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", strconv.FormatUint(uint64(intId), 10)).Msg("Media Type deleted")
}
