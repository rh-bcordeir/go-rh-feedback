package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/brunocordeiro180/go-rh-feedback/docs"
	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/server"
	"github.com/joho/godotenv"
)

// @title Swagger Feedback API
// @version 1.0
// @description This is API for avaliate candidates.
// @termsOfService http://swagger.io/terms/

// @contact.name Bruno
// @contact.email brunocordeiro@redhat.com

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := server.NewServer()

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
