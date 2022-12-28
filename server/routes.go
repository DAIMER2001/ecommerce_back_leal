package server

func (s Server) registerRoutes() {
	productRouter := s.engine.Group("api/v1/product")
	productRouter.Get("/", s.FindAllProducts)
	productRouter.Post("/", s.CreateProduct)
	productRouter.Put("/", s.UpdateProduct)
	productRouter.Delete("/", s.DeleteProduct)

	clientRouter := s.engine.Group("api/v1/client")
	clientRouter.Post("/auth", s.AuthClient)
	clientRouter.Get("/:id", s.FindAllClientById)
	clientRouter.Get("/", s.FindAllClients)
	clientRouter.Post("/", s.CreateClient)
	clientRouter.Put("/", s.UpdateClient)
	clientRouter.Delete("/", s.DeleteClient)
}
