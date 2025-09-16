package server

import (
	"encoding/json"
	"net/http"

	"github.com/brunocordeiro180/go-rh-feedback/internal/dto"
	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
)

type CandidateHandler struct {
	CandidateDB *database.CandidateDB
}

func NewCandidateHandler(db *database.CandidateDB) *CandidateHandler {
	return &CandidateHandler{CandidateDB: db}
}

func (c *CandidateHandler) CreateCandidate(w http.ResponseWriter, r *http.Request) {
	var candidateDTO dto.CandidateDTO
	if err := json.NewDecoder(r.Body).Decode(&candidateDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	candidate := candidateDTO.ToEntity()

	if err := c.CandidateDB.CreateCandidate(candidate, candidateDTO.Position); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(candidateDTO)
}

func (c *CandidateHandler) GetAllCandidates(w http.ResponseWriter, r *http.Request) {
	candidates, err := c.CandidateDB.GetAllCandidates()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(candidates)
}

func (c *CandidateHandler) UpdateCandidate(w http.ResponseWriter, r *http.Request) {
	var candidateDTO dto.CandidateDTO
	candidateID := r.PathValue("id")

	if err := json.NewDecoder(r.Body).Decode(&candidateDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	candidate := candidateDTO.ToEntityWithID(candidateID)

	if err := c.CandidateDB.UpdateCandidate(candidate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(candidateDTO)
}

func (c *CandidateHandler) DeleteCandidate(w http.ResponseWriter, r *http.Request) {
	candidateID := r.PathValue("id")

	if err := c.CandidateDB.DeleteCandidate(candidateID); err != nil {
		WriteHttpError(w, "error deleting candidate", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
