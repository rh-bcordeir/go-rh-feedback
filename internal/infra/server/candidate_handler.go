package server

import (
	"encoding/json"
	"net/http"

	"github.com/brunocordeiro180/go-rh-feedback/internal/dto"
	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
)

type CandidateHandler struct {
	CandidateDB *database.CandidateDB
}

func NewCandidateHandler(db *database.CandidateDB) *CandidateHandler {
	return &CandidateHandler{
		CandidateDB: db,
	}
}

func (c *CandidateHandler) CreateCandidate(w http.ResponseWriter, r *http.Request) {
	var candidateDTO dto.CandidateDTO
	err := json.NewDecoder(r.Body).Decode(&candidateDTO)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	candidate, err := entity.NewCandidate(candidateDTO.Name, candidateDTO.Email,
		candidateDTO.Phone, candidateDTO.Position)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.CandidateDB.CreateCandidate(candidate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(candidateDTO)
}
