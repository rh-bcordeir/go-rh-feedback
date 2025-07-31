package server

import (
	"net/http"
	"os"
	"time"

	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	db *mongo.Client
}

func NewServer() *http.Server {
	NewServer := &Server{
		db: database.NewMongoConnection(),
	}

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

	userHandler := NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.WithValue("jwtAuth", NewJWTAuth()))

	r.Post("/users/sign_in", userHandler.SignIn)

	return r
}

func NewJWTAuth() *jwtauth.JWTAuth {
	jwtSecret := os.Getenv("JWT_SECRET")
	return jwtauth.New("HS256", []byte(jwtSecret), nil)
}
