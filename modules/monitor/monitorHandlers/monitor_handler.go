package monitorHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soraritt/kawaii-shop-tutorial/config"
	"github.com/soraritt/kawaii-shop-tutorial/modules/entities"
	"github.com/soraritt/kawaii-shop-tutorial/modules/monitor"
)

type IMonitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHandler struct {
	config config.IConfig
}

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
