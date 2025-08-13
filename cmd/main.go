package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/server"
	"github.com/joho/godotenv"
)

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
