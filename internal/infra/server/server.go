package server

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/brunocordeiro180/go-rh-feedback/internal/dto"
	"github.com/brunocordeiro180/go-rh-feedback/internal/entity"
	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
	jwtPkg "github.com/brunocordeiro180/go-rh-feedback/pkg/jwt_pkg"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
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
		&entity.Position{}, &entity.Stage{}, &entity.HiringProcess{})

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
	positionDB := database.NewPositionDB(s.DB)
	stageDB := database.NewStageDB(s.DB)
	hiringProcessDB := database.NewHiringProcessDB(s.DB)

	userHandler := NewUserHandler(userDB)
	candidateHandler := NewCandidateHandler(candidateDB)
	feedbackHandler := NewFeedbackHandler(feedbackDB)
	positionHandler := NewPositionHandler(positionDB)
	stageHandler := NewStageHandler(stageDB)
	hiringProcessHandler := NewHiringProcessHandler(hiringProcessDB)

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
		r.With(jwtPkg.RequireRole("interviewer")).Delete("/{id}", candidateHandler.DeleteCandidate)
		r.With(jwtPkg.RequireRole("interviewer")).Patch("/{id}", candidateHandler.UpdateCandidate)
		r.Get("/", candidateHandler.GetAllCandidates)
	})

	r.Route("/feedbacks", func(r chi.Router) {
		r.Use(jwtauth.Verifier(TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.With(jwtPkg.RequireRole("interviewer")).Post("/", feedbackHandler.CreateFeedback)
		r.With(jwtPkg.RequireRole("interviewer")).Get("/", feedbackHandler.GetAllFeedbacks)
		r.With(jwtPkg.RequireRole("interviewer")).Delete("/{id}", feedbackHandler.DeleteFeedback)
	})

	r.Route("/positions", func(r chi.Router) {
		r.Use(jwtauth.Verifier(TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.With(jwtPkg.RequireRole("interviewer")).Post("/", positionHandler.CreatePosition)
		r.With(jwtPkg.RequireRole("interviewer")).Patch("/{id}", positionHandler.UpdatePosition)
		r.With(jwtPkg.RequireRole("interviewer")).Delete("/{id}", positionHandler.DeletePosition)
		r.Get("/", positionHandler.GetAllPositions)
	})

	r.Route("/stages", func(r chi.Router) {
		r.Use(jwtauth.Verifier(TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.With(jwtPkg.RequireRole("interviewer")).Post("/", stageHandler.CreateStage)
		r.With(jwtPkg.RequireRole("interviewer")).Delete("/{id}", stageHandler.DeleteStage)
		r.Get("/", stageHandler.GetAllStages)
	})

	r.Route("/hiring_processes", func(r chi.Router) {
		r.Use(jwtauth.Verifier(TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.With(jwtPkg.RequireRole("interviewer")).Post("/", hiringProcessHandler.CreateHiringProcess)
	})

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json"),
	))

	return r
}

func WriteHttpError(w http.ResponseWriter, errMsg string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&dto.GenericMessageDTO{Message: errMsg})
}
