package server

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/gorm"

	jwtPkg "github.com/brunocordeiro180/go-rh-feedback/pkg/jwt_pkg"
)

var (
	TokenAuth *jwtauth.JWTAuth
)

type Server struct {
	DB *gorm.DB
}

func NewJWTAuth() *jwtauth.JWTAuth {
	jwtSecret := os.Getenv("JWT_SECRET")
	return jwtauth.New("HS256", []byte(jwtSecret), nil)
}

func NewServer() *http.Server {

	db := database.NewPostgresConnection()

	// Auto-migrate entities
	_ = db.AutoMigrate(&entity.User{}, &entity.Candidate{}, &entity.Feedback{},
		&entity.Position{}, &entity.Stage{}, &entity.CandidatePosition{})

	NewServer := &Server{
		DB: db,
	}

	TokenAuth = NewJWTAuth()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) RegisterRoutes() http.Handler {

	userDB := database.NewUserDB(s.DB)
	candidateDB := database.NewCandidateDB(s.DB)
	feedbackDB := database.NewFeedbackDB(s.DB)

	userHandler := NewUserHandler(userDB)
	candidateHandler := NewCandidateHandler(candidateDB)
	feedbackHandler := NewFeedbackHandler(feedbackDB)

	expiresStr := os.Getenv("JWT_EXPIRESIN")
	expiresInt, _ := strconv.Atoi(expiresStr)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.WithValue("jwtAuth", TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", expiresInt))

	r.Post("/users/login", userHandler.SignIn)
	r.Post("/users/sign_up", userHandler.CreateUser)

	r.Route("/candidates", func(r chi.Router) {
		r.Use(jwtauth.Verifier(TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.With(jwtPkg.RequireRole("interviewer")).Post("/", candidateHandler.CreateCandidate)
		r.Get("/", candidateHandler.GetAllCandidates)
	})

	r.Route("/feedbacks", func(r chi.Router) {
		r.Use(jwtauth.Verifier(TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.With(jwtPkg.RequireRole("interviewer")).Post("/", feedbackHandler.CreateFeedback)
	})

	return r
}
