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

// Sign In godoc
// @Summary      Authenticate user and return JWT
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request  body      dto.SignInRequest  true  "sign in request"
// @Success      200      {object}  dto.GetJWTOutput
// @Failure      400      {object}  dto.GenericMessageDTO
// @Failure      401      {object}  dto.GenericMessageDTO
// @Router       /users/login [post]
func (u *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {

	var input dto.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		WriteHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.UserDB.FindByEmail(input.Email)
	if err != nil || !user.ValidatePassword(input.Password) {
		WriteHttpError(w, "invalid email or password", http.StatusUnauthorized)
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

// Create User godoc
// @Summary      Create a new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateUserDTO  true  "create user request"
// @Success      200      {object}  dto.GenericMessageDTO
// @Failure      400      {object}  dto.GenericMessageDTO
// @Failure      500      {object}  dto.GenericMessageDTO
// @Router       /users/sign_up [post]
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	reqID := middleware.GetReqID(r.Context())
	log.Printf("CreateUser called. reqID=%s", reqID)

	var createUserDTO dto.CreateUserDTO
	err := json.NewDecoder(r.Body).Decode(&createUserDTO)
	if err != nil {
		WriteHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: validate password
	// For now the email will be already verified but in the future it will send
	// an email confirmation
	user := createUserDTO.ToEntity()

	if err := user.ValidateEmail(user.Email); err != nil {
		WriteHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := user.HashPassword(); err != nil {
		WriteHttpError(w, "failed to hash password", http.StatusInternalServerError)
		return
	}

	if err := u.UserDB.SaveUser(user); err != nil {
		WriteHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: "User Created"})
}
