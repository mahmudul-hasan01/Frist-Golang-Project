package cmd

import (
	"back-end/config"
	"back-end/infra/db"
	"back-end/repo"
	"back-end/rest"
	"back-end/rest/handler/product"
	"back-end/rest/handler/user"
	"back-end/rest/middleware"
	"fmt"
	"os"
)

func Server() {
	cnf := config.GetConfig()
	dbcon, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo()
	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userRepo := repo.NewUserRepo(dbcon)
	userHandler := user.NewHandler(middlewares, userRepo, cnf)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start(*cnf)

}
