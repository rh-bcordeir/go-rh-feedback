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
