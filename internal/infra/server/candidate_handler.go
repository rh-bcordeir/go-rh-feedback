package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/brunocordeiro180/go-rh-feedback/internal/dto"
	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
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

	candidate := &entity.Candidate{
		Name:      candidateDTO.Name,
		Email:     candidateDTO.Email,
		Phone:     candidateDTO.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := c.CandidateDB.CreateCandidate(candidate, candidateDTO.Position); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(candidateDTO)
}
