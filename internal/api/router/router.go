package router

import (
	"github.com/ReeceRose/home-network-proxy/internal/api/controller"
	"github.com/labstack/echo/v4"
)

// Setup handles all the public/private server routes
func Setup(e *echo.Echo) {
	// General setup
	ip := controller.NewIPController()

	// Public routes

	// TODO: setup auth middleware

	// Private routes
	e.GET("/api/v1/ip/", func(c echo.Context) error { return ip.GetIP(c) })
	e.POST("/api/v1/ip/", func(c echo.Context) error { return ip.PostIP(c) })
}
