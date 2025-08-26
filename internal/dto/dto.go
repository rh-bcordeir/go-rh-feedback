package dto

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}

type GenericMessageDTO struct {
	Message string `json:"message"`
}

type CandidateDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Position uint   `json:"position"`
}

type FeedbackDTO struct {
	CandidateID string `json:"candidate_id"`
	StageID     uint   `json:"stage_id"`
	Comments    string `json:"comments"`
	Score       int    `json:"score"`
}
