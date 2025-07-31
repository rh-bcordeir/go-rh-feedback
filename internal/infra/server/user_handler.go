package server

import (
	"encoding/json"
	"net/http"

	"github.com/brunocordeiro180/go-rh-feedback/internal/dto"
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

	_ = r.Context().Value("jwtAuth").(*jwtauth.JWTAuth) //change later
	var userDTO dto.SignInRequest
	err := json.NewDecoder(r.Body).Decode(&userDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = u.UserDB.FindByEmail(userDTO.Email) //change later
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
