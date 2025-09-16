package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/brunocordeiro180/go-rh-feedback/internal/dto"
	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB *database.UserDB
}

func NewUserHandler(db *database.UserDB) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

func (u *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input dto.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.UserDB.FindByEmail(input.Email)
	if err != nil || !user.ValidatePassword(input.Password) {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	jwtAuth := r.Context().Value("jwtAuth").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExpiresIn").(int)

	claims := map[string]interface{}{
		"sub":   user.ID.String(),
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Duration(jwtExpiresIn) * time.Second).Unix(),
	}

	_, tokenString, _ := jwtAuth.Encode(claims)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&dto.GetJWTOutput{AccessToken: tokenString})
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	reqID := middleware.GetReqID(r.Context())
	log.Printf("CreateUser called. reqID=%s", reqID)

	var createUserDTO dto.CreateUserDTO
	err := json.NewDecoder(r.Body).Decode(&createUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := dto.GenericMessageDTO{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	// TODO: validate password
	// For now the email will be already verified but in the future it will send
	// an email confirmation
	user := createUserDTO.ToEntity()

	if err := user.ValidateEmail(user.Email); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	if err := user.HashPassword(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	if err := u.UserDB.SaveUser(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.GenericMessageDTO{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: "User Created"})
}
