package monitorHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soraritt/kawaii-shop-tutorial/config"
	"github.com/soraritt/kawaii-shop-tutorial/modules/entities"
	"github.com/soraritt/kawaii-shop-tutorial/modules/monitor"
)

// interface
type IMonitorHandler interface {
	//hanlder ของ fiber  , *fiber.Ctx
	HealthCheck(c *fiber.Ctx) error
}

// struct
type monitorHandler struct {
	config config.IConfig
}

// constructor
func MonitorHandler(config config.IConfig) IMonitorHandler {
	return &monitorHandler{
		config: config,
	}
}

func (h *monitorHandler) HealthCheck(c *fiber.Ctx) error {
	res := &monitor.Monitor{
		Name:    h.config.App().Name(),
		Version: h.config.App().Version(),
	}
	return entities.NewResponse(c).Success(fiber.StatusOK, res).Res()
}
