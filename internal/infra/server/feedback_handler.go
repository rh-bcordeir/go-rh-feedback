package server

import (
	"encoding/json"
	"net/http"

	"github.com/brunocordeiro180/go-rh-feedback/internal/dto"
	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
	jwtPkg "github.com/brunocordeiro180/go-rh-feedback/pkg/jwt_pkg"
)

type FeedbackHandler struct {
	FeedbackDB *database.FeedbackDB
}

func NewFeedbackHandler(db *database.FeedbackDB) *FeedbackHandler {
	return &FeedbackHandler{FeedbackDB: db}
}

func (f *FeedbackHandler) CreateFeedback(w http.ResponseWriter, r *http.Request) {
	var feedbackDTO dto.FeedbackDTO
	if err := json.NewDecoder(r.Body).Decode(&feedbackDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := jwtPkg.ExtractUserIDFromToken(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feedback := entity.NewFeedback(
		userId,
		feedbackDTO.CandidateID,
		feedbackDTO.StageID,
		feedbackDTO.Comments,
		feedbackDTO.Score,
	)

	if err := f.FeedbackDB.SaveFeedback(feedback); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(feedbackDTO)
}
