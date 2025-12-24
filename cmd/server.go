package cmd

import (
	"back-end/config"
	"back-end/rest"
	"back-end/rest/handler/product"
	"back-end/rest/handler/user"
	"back-end/rest/middleware"
)

func Server() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler(middlewares)

	rest.NewServer(cnf, productHandler, userHandler)
	// server.Start(cnf)

}
