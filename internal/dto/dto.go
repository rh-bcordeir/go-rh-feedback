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
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type FeedbackDTO struct {
	StageID         uint   `json:"stage_id"`
	HiringProcessID uint   `json:"hiring_process_id"`
	Comments        string `json:"comments"`
	Score           int    `json:"score"`
}

type HiringProcessDTO struct {
	CandidateID string `json:"candidate_id"`
	PositionID  uint   `json:"position_id"`
}

type PositionDTO struct {
	Title string `json:"title"`
}

type StageDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
