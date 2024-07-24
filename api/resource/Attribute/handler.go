package attribute

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

type AttributeAPI struct {
	logger     *zerolog.Logger
	repository *AttributeRepository
}

func NewAttributeAPI(logger *zerolog.Logger, db *gorm.DB) *AttributeAPI {
	return &AttributeAPI{
		logger:     logger,
		repository: NewAttributeRepository(db),
	}
}

func (api AttributeAPI) Create(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())

	form := &AttributeForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	attribute, err := NewAttribute(form.Name)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDomainRulesFailure)
		return
	}

	addAttribute, err := api.repository.Create(attribute)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}

	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", strconv.FormatUint(uint64(addAttribute.ID), 10)).Msg("New Attribute was added")
	w.WriteHeader(http.StatusCreated)

}

func (api AttributeAPI) Read(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	intId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	attribute, err := api.repository.Read(uint(intId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	if err := json.NewEncoder(w).Encode(attribute.ToDTO()); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

func (api AttributeAPI) Update(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	intId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	form := &AttributeForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONDecodeFailure)
		return
	}

	attribute, err := NewAttribute(form.Name)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDomainRulesFailure)
		return

	}
	attribute.ID = uint(intId)
	rows, err := api.repository.Update(attribute)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataUpdateFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", strconv.FormatUint(uint64(attribute.ID), 10)).Msg("Attribute was updated")
}

func (api AttributeAPI) Delete(w http.ResponseWriter, r *http.Request) {
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

	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", strconv.FormatUint(uint64(intId), 10)).Msg("Attribute was deleted")
}
