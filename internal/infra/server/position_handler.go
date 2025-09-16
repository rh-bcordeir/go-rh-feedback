package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/brunocordeiro180/go-rh-feedback/internal/dto"
	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
)

type PositionHandler struct {
	PositionDB *database.PositionDB
}

func NewPositionHandler(db *database.PositionDB) *PositionHandler {
	return &PositionHandler{PositionDB: db}
}

func (p *PositionHandler) GetAllPositions(w http.ResponseWriter, r *http.Request) {
	positions, err := p.PositionDB.GetAllPositions()
	if err != nil {
		WriteHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(positions)
}

func (p *PositionHandler) CreatePosition(w http.ResponseWriter, r *http.Request) {
	var positionDTO dto.PositionDTO
	if err := json.NewDecoder(r.Body).Decode(&positionDTO); err != nil {
		WriteHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.PositionDB.CreatePosition(positionDTO.ToEntity()); err != nil {
		WriteHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(positionDTO)
}

func (p *PositionHandler) UpdatePosition(w http.ResponseWriter, r *http.Request) {
	var positionDTO dto.PositionDTO
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		WriteHttpError(w, "invalid position ID", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&positionDTO); err != nil {
		WriteHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.PositionDB.UpdatePosition(positionDTO.ToEntityWithID(id)); err != nil {
		WriteHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(positionDTO)
}

func (p *PositionHandler) DeletePosition(w http.ResponseWriter, r *http.Request) {
	positionID := r.PathValue("id")
	id, err := strconv.ParseUint(positionID, 10, 64)

	if err != nil {
		WriteHttpError(w, "invalid position ID", http.StatusBadRequest)
		return
	}

	if err := p.PositionDB.DeletePosition(uint(id)); err != nil {
		WriteHttpError(w, "error deleting position", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
