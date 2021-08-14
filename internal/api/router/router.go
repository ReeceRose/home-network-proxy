package router

import (
	"github.com/labstack/echo/v4"
)

// Setup handles all the public/private server routes
func Setup(e *echo.Echo) {
	// General setup
	// health := controller.NewHealthController()
	// host := controller.NewHostController()

	// Public routes

	// TODO: setup auth middleware

	// Private routes
	e.GET("/api/v1/ip/", func(c echo.Context) error { return nil })
	e.POST("/api/v1/ip/", func(c echo.Context) error { return nil })
}
