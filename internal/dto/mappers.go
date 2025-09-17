package dto

import (
	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"github.com/google/uuid"
)

func (c *CandidateDTO) ToEntity() *entity.Candidate {
	return &entity.Candidate{
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
	}
}

func (c *CandidateDTO) ToEntityWithID(id string) *entity.Candidate {
	candidateUUID := uuid.MustParse(id)

	return &entity.Candidate{
		ID:    candidateUUID,
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
	}
}

func (p *PositionDTO) ToEntity() *entity.Position {
	return &entity.Position{
		Title: p.Title,
	}
}

func (p *PositionDTO) ToEntityWithID(id uint64) *entity.Position {
	return &entity.Position{
		ID:    uint(id),
		Title: p.Title,
	}
}

func (s *StageDTO) ToEntity() *entity.Stage {
	return &entity.Stage{
		Title:       s.Title,
		Description: s.Description,
	}
}

func (u *CreateUserDTO) ToEntity() *entity.User {
	return &entity.User{
		Name:          u.Name,
		Email:         u.Email,
		Password:      u.Password,
		Role:          entity.INTERVIEWER,
		EmailVerified: true,
	}
}

func (h *HiringProcessDTO) ToEntity() *entity.HiringProcess {
	candidateID := uuid.MustParse(h.CandidateID)

	return &entity.HiringProcess{
		CandidateID: candidateID,
		PositionID:  h.PositionID,
	}
}
