package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/brunocordeiro180/go-rh-feedback/internal/dto"
	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
)

type StageHandler struct {
	StageDB *database.StageDB
}

func NewStageHandler(stageDB *database.StageDB) *StageHandler {
	return &StageHandler{StageDB: stageDB}
}

// Get All Stages godoc
// @Summary      List all stages
// @Tags         Stages
// @Produce      json
// @Success      200  {array}   dto.StageDTO
// @Failure      400  {object}  dto.GenericMessageDTO
// @Router       /stages [get]
// @Security     ApiKeyAuth
func (s *StageHandler) GetAllStages(w http.ResponseWriter, r *http.Request) {
	AddPrometheusMetrics(r)

	stages, err := s.StageDB.GetAllStages()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stages)
}

// Create Stage godoc
// @Summary      Create stage
// @Tags         Stages
// @Accept       json
// @Produce      json
// @Param        request  body      dto.StageDTO  true  "stage request"
// @Success      201      {object}  dto.StageDTO
// @Failure      400      {object}  dto.GenericMessageDTO
// @Router       /stages [post]
// @Security     ApiKeyAuth
func (s *StageHandler) CreateStage(w http.ResponseWriter, r *http.Request) {
	AddPrometheusMetrics(r)

	var stageDTO dto.StageDTO
	if err := json.NewDecoder(r.Body).Decode(&stageDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.StageDB.CreateStage(stageDTO.ToEntity()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(stageDTO)
}

// Delete Stage godoc
// @Summary      Delete stage
// @Tags         Stages
// @Param        id   path  string  true  "Stage ID"
// @Success      204
// @Failure      400  {object}  dto.GenericMessageDTO
// @Router       /stages/{id} [delete]
// @Security     ApiKeyAuth
func (s *StageHandler) DeleteStage(w http.ResponseWriter, r *http.Request) {
	AddPrometheusMetrics(r)

	stageID := r.PathValue("id")
	id, err := strconv.ParseUint(stageID, 10, 64)
	if err != nil {
		WriteHttpError(w, "Invalid stage ID", http.StatusBadRequest)
		return
	}

	if err := s.StageDB.DeleteStage(uint(id)); err != nil {
		WriteHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
