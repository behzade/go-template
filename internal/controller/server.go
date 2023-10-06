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

	router.GET("openapi.json", func(c echo.Context) error {
		swagger, err := GetSwagger()
		swagger.Servers = nil
		if err != nil {
			return err
		}

		return c.JSON(200, swagger)
	})

	RegisterHandlers(router, server)

	return router
}
