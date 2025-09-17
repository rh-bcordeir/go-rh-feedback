package server

import (
	"encoding/json"
	"net/http"

	"github.com/brunocordeiro180/go-rh-feedback/internal/dto"
	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
)

type HiringProcessHandler struct {
	HiringProcessDB *database.HiringProcessDB
}

func NewHiringProcessHandler(db *database.HiringProcessDB) *HiringProcessHandler {
	return &HiringProcessHandler{
		HiringProcessDB: db,
	}
}

// Create HiringProcess godoc
// @Summary      Create Hiring Process
// @Tags         HiringProcesses
// @Accept       json
// @Produce      json
// @Param        request  body      dto.HiringProcessDTO  true  "hiring_process request"
// @Success      201      {object}  entity.HiringProcess
// @Failure      400      {object}  dto.GenericMessageDTO
// @Router       /hiring_processes [post]
// @Security     ApiKeyAuth
func (h *HiringProcessHandler) CreateHiringProcess(w http.ResponseWriter, r *http.Request) {
	var hiringProcessDTO dto.HiringProcessDTO
	if err := json.NewDecoder(r.Body).Decode(&hiringProcessDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hiringProcess := hiringProcessDTO.ToEntity()

	if err := h.HiringProcessDB.SaveHiringProcess(hiringProcess); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(hiringProcessDTO)
}
