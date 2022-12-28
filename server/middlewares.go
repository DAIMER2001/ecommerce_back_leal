package server

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (s *Server) registerMiddlewares() {
	s.engine.Use(recover.New())

	s.engine.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} - ${body}\n",
	}))

	config := cors.Config{
		AllowOrigins:     s.configApp.AllowOrigins,
		AllowMethods:     "HEAD,GET,POST,PUT,DELETE,PATCH",
		MaxAge:           120,
		AllowCredentials: true,
	}
	s.engine.Use(cors.New(config))

}
