package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soraritt/kawaii-shop-tutorial/modules/middlewares/middlewaresHandlers"
	"github.com/soraritt/kawaii-shop-tutorial/modules/middlewares/middlewaresRepositories"
	"github.com/soraritt/kawaii-shop-tutorial/modules/middlewares/middlewaresUsecases"
	"github.com/soraritt/kawaii-shop-tutorial/modules/monitor/monitorHandlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.database)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	h := middlewaresHandlers.MiddlewaresHandler(s.config, usecase)

	return h
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.config)

	m.r.Get("/", handler.HealthCheck)

}
