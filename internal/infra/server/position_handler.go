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

// Get All Positions godoc
// @Summary      List all positions
// @Tags         Positions
// @Produce      json
// @Success      200  {array}   dto.PositionDTO
// @Failure      400  {object}  dto.GenericMessageDTO
// @Router       /positions [get]
// @Security     ApiKeyAuth
func (p *PositionHandler) GetAllPositions(w http.ResponseWriter, r *http.Request) {
	AddPrometheusMetrics(r)

	positions, err := p.PositionDB.GetAllPositions()
	if err != nil {
		WriteHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(positions)
}

// Create Position godoc
// @Summary      Create position
// @Tags         Positions
// @Accept       json
// @Produce      json
// @Param        request  body      dto.PositionDTO  true  "position request"
// @Success      201      {object}  dto.PositionDTO
// @Failure      400      {object}  dto.GenericMessageDTO
// @Router       /positions [post]
// @Security     ApiKeyAuth
func (p *PositionHandler) CreatePosition(w http.ResponseWriter, r *http.Request) {
	AddPrometheusMetrics(r)

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

// Update Position godoc
// @Summary      Update position
// @Tags         Positions
// @Accept       json
// @Produce      json
// @Param        id       path      string          true  "Position ID"
// @Param        request  body      dto.PositionDTO true  "position request"
// @Success      200      {object}  dto.PositionDTO
// @Failure      400      {object}  dto.GenericMessageDTO
// @Router       /positions/{id} [put]
// @Security     ApiKeyAuth
func (p *PositionHandler) UpdatePosition(w http.ResponseWriter, r *http.Request) {
	AddPrometheusMetrics(r)

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

// Delete Position godoc
// @Summary      Delete position
// @Tags         Positions
// @Param        id   path  string  true  "Position ID"
// @Success      204
// @Failure      400  {object}  dto.GenericMessageDTO
// @Router       /positions/{id} [delete]
// @Security     ApiKeyAuth
func (p *PositionHandler) DeletePosition(w http.ResponseWriter, r *http.Request) {
	AddPrometheusMetrics(r)

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
