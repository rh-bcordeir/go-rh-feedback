package main

import (
	"fmt"
	"net/http"

	"github.com/brunocordeiro180/go-rh-feedback/internal/infra/server"
)

func main() {
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

}
