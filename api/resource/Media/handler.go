package media

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

type MediaApi struct {
	logger     *zerolog.Logger
	repository *MediaRepository
}

func NewMediaAPI(logger *zerolog.Logger, database *gorm.DB) *MediaApi {
	return &MediaApi{
		logger:     logger,
		repository: NewMediaRepository(database),
	}
}

func (api MediaApi) Create(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	form := &MediaForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}
	media, err := NewMedia(form.Name, form.MediaTypeID)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDomainRulesFailure)
		return
	}

	addMedia, err := api.repository.Create(media)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}

	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", strconv.FormatUint(uint64(addMedia.ID), 10)).Msg("New Media was added")
	w.WriteHeader(http.StatusCreated)
}

func (api MediaApi) Read(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	idStr := chi.URLParam(r, "id")
	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", idStr).Msg("Captured URL parameter")

	intId, err := strconv.Atoi(idStr)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	media, err := api.repository.Read(uint(intId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
		}
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}
	api.logger.Info().Msg(media.MediaType.Name)
	dto := media.ToDTO()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

func (api MediaApi) Update(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	intId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	form := &MediaForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONDecodeFailure)
		return
	}

	media, err := NewMedia(form.Name, form.MediaTypeID)
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

	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", strconv.FormatUint(uint64(intId), 10)).Msg("Updated Media")

}

func (api MediaApi) Delete(w http.ResponseWriter, r *http.Request) {
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
		e.ServerError(w, e.RespDBDataUpdateFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", strconv.FormatUint(uint64(intId), 10)).Msg("Deleted Media")
}
