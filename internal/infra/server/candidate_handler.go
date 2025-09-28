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

// Create Candidate godoc
// @Summary      Create candidate
// @Tags         Candidates
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CandidateDTO  true  "product request"
// @Success      201
// @Failure      400         {object}  dto.GenericMessageDTO
// @Router       /candidates [post]
// @Security ApiKeyAuth
func (c *CandidateHandler) CreateCandidate(w http.ResponseWriter, r *http.Request) {

	var candidateDTO dto.CandidateDTO
	if err := json.NewDecoder(r.Body).Decode(&candidateDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	candidate := candidateDTO.ToEntity()

	if err := c.CandidateDB.CreateCandidate(candidate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(candidateDTO)
}

// Get All Candidates godoc
// @Summary      List all candidates
// @Tags         Candidates
// @Produce      json
// @Success      200  {array}   dto.CandidateDTO
// @Failure      400  {object}  dto.GenericMessageDTO
// @Router       /candidates [get]
// @Security     ApiKeyAuth
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

// Update Candidate godoc
// @Summary      Update candidate
// @Tags         Candidates
// @Accept       json
// @Produce      json
// @Param        id        path      string           true  "Candidate ID"
// @Param        request   body      dto.CandidateDTO true  "candidate request"
// @Success      200       {object}  dto.CandidateDTO
// @Failure      400       {object}  dto.GenericMessageDTO
// @Router       /candidates/{id} [put]
// @Security     ApiKeyAuth
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

// Delete Candidate godoc
// @Summary      Delete candidate
// @Tags         Candidates
// @Param        id   path  string  true  "Candidate ID"
// @Success      204
// @Failure      400  {object}  dto.GenericMessageDTO
// @Router       /candidates/{id} [delete]
// @Security     ApiKeyAuth
func (c *CandidateHandler) DeleteCandidate(w http.ResponseWriter, r *http.Request) {
	candidateID := r.PathValue("id")

	if err := c.CandidateDB.DeleteCandidate(candidateID); err != nil {
		WriteHttpError(w, "error deleting candidate", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
