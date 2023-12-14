package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soraritt/kawaii-shop-tutorial/modules/middlewares/middlewaresHandlers"
	"github.com/soraritt/kawaii-shop-tutorial/modules/middlewares/middlewaresRepositories"
	"github.com/soraritt/kawaii-shop-tutorial/modules/middlewares/middlewaresUsecases"
	"github.com/soraritt/kawaii-shop-tutorial/modules/monitor/monitorHandlers"
	"github.com/soraritt/kawaii-shop-tutorial/modules/users/usersHandlers"
	"github.com/soraritt/kawaii-shop-tutorial/modules/users/usersRepositories"
	"github.com/soraritt/kawaii-shop-tutorial/modules/users/usersUsecases"
)

type IModuleFactory interface {
	MonitorModule()
	UsersModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

// constructor
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
	m.r.Get("/helloworld", func(c *fiber.Ctx) error {
		return c.SendString("hello world 🌈")
	})
	m.r.Get("/info", func(c *fiber.Ctx) error { // JSON
		return c.JSON(fiber.Map{
			"msg":     "hello world 🚀",
			"go":      "fiber 🥦",
			"boolean": true,
			"number":  1234,
		})
	})
}

func (m *moduleFactory) UsersModule() {
	repository := usersRepositories.UsersRepository(m.s.database)
	usecase := usersUsecases.UsersUsecase(m.s.config, repository)
	handler := usersHandlers.UsersHandler(m.s.config, usecase)

	router := m.r.Group("/users")
	//  /users/signup
	router.Post("/signup", handler.SignUpCustomer)
	router.Post("/signin", handler.SignIn)
	router.Post("/refresh", handler.RefreshPassport)
	router.Post("/signout", handler.SignOut)
	router.Post("/signup-admin", handler.SignOut)
	router.Get("/:user_id", m.mid.JwtAuth(), m.mid.ParamsCheck(), handler.GetUserProfile)
	router.Get("/admin/secret", m.mid.JwtAuth(), m.mid.Authorize(2, 1), handler.GenerateAdminToken)

	//Initail admin 1 คน ใน database (Insert ใน sql)
	//Generate Admin key Token
	//ทุกครั้งที่สมัคร admin เพิ่มให้ส่ง  Admin Token  มาด้วยทุกครั้งผ่าน  middleware

}
