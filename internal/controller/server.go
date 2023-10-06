package controller

import "github.com/labstack/echo/v4"

type Server struct{}

func New() *echo.Echo {
	server := Server{}
	router := echo.New()

	RegisterHandlers(router, &server)

	return router
}

// (GET /)
func (server *Server) Index(ctx echo.Context) error {
	return indexHandler(ctx)
}

// (GET /healthz)
func (server *Server) Healthz(ctx echo.Context) error {
	return checkHealthHandler(ctx)
}
