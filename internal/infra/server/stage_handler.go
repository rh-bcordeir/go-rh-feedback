package server

import (
	"encoding/json"
	"net/http"

	"github.com/brunocordeiro180/go-rh-feedback/internal/dto"
	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
)

type StageHandler struct {
	StageDB *database.StageDB
}

func NewStageHandler(stageDB *database.StageDB) *StageHandler {
	return &StageHandler{StageDB: stageDB}
}

func (s *StageHandler) GetAllStages(w http.ResponseWriter, r *http.Request) {
	stages, err := s.StageDB.GetAllStages()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stages)
}

func (s *StageHandler) CreateStage(w http.ResponseWriter, r *http.Request) {
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
