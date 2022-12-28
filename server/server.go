package server

import (
	"ecommerce/config"
	"ecommerce/features"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	configApp      config.Config
	engine         *fiber.App
	clientService  features.ClientFeatures
	productService features.ProductFeatures
}

func New(cfg config.Config,
	clientService features.ClientFeatures,
	productService features.ProductFeatures,
) Server {
	svr := Server{
		configApp:      cfg,
		clientService:  clientService,
		productService: productService,
	}

	svr.engine = fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	svr.registerMiddlewares()
	svr.registerRoutes()
	return svr
}

func (s Server) Run() error {
	return s.engine.Listen(s.configApp.HttpAddress)
}
