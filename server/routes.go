package server

func (s Server) registerRoutes() {
	productRouter := s.engine.Group("api/v1/product")
	productRouter.Get("/", s.FindAllProducts)
	productRouter.Post("/", s.CreateProduct)
	productRouter.Put("/", s.UpdateProduct)
	productRouter.Delete("/", s.DeleteProduct)

	clientRouter := s.engine.Group("api/v1/client")
	clientRouter.Post("/auth", s.AuthClient)
	clientRouter.Get("/points/:name", s.GetClientByName)
	clientRouter.Get("/:id", s.GetClientById)
	clientRouter.Get("/", s.FindAllClients)
	clientRouter.Post("/", s.CreateClient)
	clientRouter.Put("/points", s.UpdatePointsClient)
	clientRouter.Put("/redime-points", s.RedimeProduct)

	clientRouter.Delete("/", s.DeleteClient)
}
