package cmd

import (
	"back-end/config"
	"back-end/repo"
	"back-end/rest"
	"back-end/rest/handler/product"
	"back-end/rest/handler/user"
	"back-end/rest/middleware"
)

func Server() {
	cnf := config.GetConfig()
	productRepo := repo.NewProductRepo()
	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userRepo := repo.NewUserRepo()
	userHandler := user.NewHandler(middlewares, userRepo, cnf)

	rest.NewServer(cnf, productHandler, userHandler)
	// server.Start(cnf)

}
