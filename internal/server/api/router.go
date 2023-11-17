package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// TODO: Implment user authentication
	// r.Use(authMiddleware)

	// Version 1 API Routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/templates", s.templatesHandlerV1)
	})

	return r
}
