package servers

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/soraritt/kawaii-shop-tutorial/config"
)

type IServer interface {
	Start()
}

type server struct {
	app      *fiber.App
	config   config.IConfig
	database *sqlx.DB
}

func NewServer(config config.IConfig, database *sqlx.DB) IServer {

	return &server{
		config:   config,
		database: database,
		app: fiber.New(fiber.Config{
			AppName:      config.App().Name(),
			BodyLimit:    config.App().BodyLimit(),
			ReadTimeout:  config.App().ReadTimeout(),
			WriteTimeout: config.App().WriteTimeout(),
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
	}

}

func (s *server) Start() {
	//Middlewares
	middlewares := InitMiddlewares(s)
	s.app.Use(middlewares.Logger())
	s.app.Use(middlewares.Cors())

	//Modules
	v1 := s.app.Group("v1")
	modules := InitModule(v1, s, middlewares)
	modules.MonitorModule()

	s.app.Use(middlewares.RouterCheck())

	//Graceful shutdown
	c := make(chan os.Signal, 1)   // เปิด 1 channel  รับสัญญาณ
	signal.Notify(c, os.Interrupt) //check interrupt
	go func() {
		_ = <-c
		log.Println("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	//Listen to host:port
	log.Printf("server is starting on %v", s.config.App().Url())
	s.app.Listen(s.config.App().Url())

}
