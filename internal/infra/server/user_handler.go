package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/brunocordeiro180/go-rh-feedback/internal/dto"
	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
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
	jwtAuth := r.Context().Value("jwtAuth").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExpiresIn").(int)
	var userDTO dto.SignInRequest
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := u.UserDB.FindByEmail(userDTO.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := dto.GenericMessageDTO{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Println(user)

	if !user.ValidatePassword(userDTO.Password) || user.EmailVerified.IsZero() {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwtAuth.Encode(map[string]interface{}{
		"sub":   user.ID.String(),
		"exp":   time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
		"roles": []string{string(user.Role)},
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

func (u *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var createUserDTO dto.CreateUserDTO
	err := json.NewDecoder(r.Body).Decode(&createUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := dto.GenericMessageDTO{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	// TODO: validate password
	user, err := entity.NewUser(createUserDTO.Name, createUserDTO.Email, createUserDTO.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.GenericMessageDTO{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	if err = user.ValidateEmail(user.Email); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.GenericMessageDTO{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	// For now the email will be already verified but in the future it will send
	// an email confirmation
	user.EmailVerified = time.Now()

	err = u.UserDB.SaveUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.GenericMessageDTO{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		&dto.GenericMessageDTO{
			Message: "User Created",
		},
	)
}
