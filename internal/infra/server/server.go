package server

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	TokenAuth *jwtauth.JWTAuth
)

type Server struct {
	db *mongo.Client
}

func NewServer() *http.Server {
	NewServer := &Server{
		db: database.NewMongoConnection(),
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

	userDB := database.NewUserDB(s.db)
	candidateDB := database.NewCandidateDB(s.db)

	userHandler := NewUserHandler(userDB)
	candidateHandler := NewCandidateHandler(candidateDB)

	expiresStr := os.Getenv("JWT_EXPIRESIN")
	expiresInt, _ := strconv.Atoi(expiresStr)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.WithValue("jwtAuth", TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", expiresInt))

	r.Post("/users/login", userHandler.SignIn)
	r.Post("/users/sign_up", userHandler.SignUp)

	r.Route("/candidates", func(r chi.Router) {
		r.Use(jwtauth.Verifier(TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", candidateHandler.CreateCandidate)
	})

	return r
}

func NewJWTAuth() *jwtauth.JWTAuth {
	jwtSecret := os.Getenv("JWT_SECRET")
	return jwtauth.New("HS256", []byte(jwtSecret), nil)
}
