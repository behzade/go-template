package controller

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Server struct{}

// (GET /)
func (server *Server) Index(ctx context.Context, request IndexRequestObject) (IndexResponseObject, error) {
	return Index200JSONResponse{"go template server"}, nil
}

// (GET /healthz)
func (server *Server) Healthz(ctx context.Context, request HealthzRequestObject) (HealthzResponseObject, error) {
	return Healthz200JSONResponse{"ok"}, nil
}

func New() *echo.Echo {
	server := NewStrictHandler(&Server{}, nil)
	router := echo.New()

	RegisterHandlers(router, server)

	return router
}
