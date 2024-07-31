package role

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

type RoleAPI struct {
	logger *zerolog.Logger
	repo   RoleRepository
}

func NewRoleAPI(logger *zerolog.Logger, db *gorm.DB) *RoleAPI {
	return &RoleAPI{
		logger: logger,
		repo:   *NewRoleRepository(db),
	}
}

func (api RoleAPI) Create(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	form := RoleForm{}
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	role, err := NewRole(form.Name, form.Description)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDomainRulesFailure)
		return
	}

	createdRole, err := api.repo.Create(role)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}
	api.logger.Info().Str(log.KeyReqID, reqID).Str("id: ", strconv.FormatUint(uint64(createdRole.ID), 10)).Msg("New Role inserted")
	w.WriteHeader(http.StatusCreated)
}
func (api *RoleAPI) Read(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	idStr := chi.URLParam(r, "id")

	intId, err := strconv.Atoi(idStr)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	role, err := api.repo.Read(uint(intId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
		}
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}
	dto := role.ToDTO()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

func (api *RoleAPI) Update(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	intId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	form := &RoleForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONDecodeFailure)
		return
	}

	role, err := NewRole(form.Name, form.Description)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDomainRulesFailure)
		return

	}
	role.ID = uint(intId)
	rows, err := api.repo.Update(role)
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataUpdateFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	api.logger.Info().Str(log.KeyReqID, reqID).Str("id", strconv.FormatUint(uint64(intId), 10)).Msg("Updated Role")

}

func (api *RoleAPI) Delete(w http.ResponseWriter, r *http.Request) {
	reqID := ctx.RequestID(r.Context())
	intId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.logger.Error().Str(log.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}
	rows, err := api.repo.Delete(uint(intId))
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
