package controller

import (
	"github.com/ReeceRose/home-network-proxy/internal/repository"
	"github.com/ReeceRose/home-network-proxy/internal/service"
	"github.com/ReeceRose/home-network-proxy/internal/types"
	"github.com/labstack/echo/v4"
)

// IPController provides a IP service to interact with
type IPController struct {
	ipService service.IIPService
}

// NewIPController returns a new IPController with the service/repository initialized
func NewIPController() *IPController {
	return &IPController{
		ipService: service.NewIPService(repository.NewIPRepository()),
	}
}

// GetAllIP returns all reported external IPs
func (controller *IPController) GetAllIP(c echo.Context) error {
	res := controller.ipService.GetAllIP(
		c.Response().Header().Get("X-Request-ID"),
	)
	return c.JSON(res.StatusCode, res)
}

// PostIP inserts new reported IPs and updates existing reported IPs
func (controller *IPController) PostIP(c echo.Context) error {
	data := new(types.IP)

	if err := c.Bind(data); err != nil {
		return c.JSON(400, types.IPResponse{
			StatusCode: 400,
			Error:      "failed to bind IP",
			Success:    false,
			Data:       []types.IP{},
		})
	}

	res := controller.ipService.InsertIP(
		c.Response().Header().Get("X-Request-ID"),
		data,
	)

	return c.JSON(res.StatusCode, res)
}
