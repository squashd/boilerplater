package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

var port = 8080

type Server struct {
	port int
}

func NewServer() *http.Server {

	// Loads the super duper secret API Key
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	NewServer := &Server{
		port: port,
	}

	// Requests can take a while so I've tried to be as lenient as possible
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  5 * time.Minute,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 2 * time.Minute,
	}

	return server
}
