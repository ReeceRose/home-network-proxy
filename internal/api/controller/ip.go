package controller

import (
	"github.com/labstack/echo/v4"
)

// IPController provides a IP service to interact with
type IPController struct {
}

// NewIPController returns a new IPController with the service/repository initialized
func NewIPController() *IPController {
	return &IPController{}
}

// GetIP returns all reported external IPs
func (controller *IPController) GetIP(c echo.Context) error {
	return c.JSON(200, "")
}

// PostIP inserts new reported IPs and updates existing reported IPs
func (controller *IPController) PostIP(c echo.Context) error {
	return c.JSON(200, "")
}
