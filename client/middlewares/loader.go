package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Load registers the API middlewares
func Load(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return idGenerator()
		},
	}))
}

func idGenerator() (id string) {
	return uuid.New().String()
}
