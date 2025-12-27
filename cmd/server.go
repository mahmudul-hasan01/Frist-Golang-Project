package cmd

import (
	"back-end/config"
	"back-end/infra/db"
	"back-end/repo"
	"back-end/rest"
	prdHandler "back-end/rest/handler/product"
	usrHandler "back-end/rest/handler/user"
	"back-end/rest/middleware"
	"fmt"
	"os"
)

func Server() {
	cnf := config.GetConfig()
	dbcon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err := db.MigrateDB(dbcon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbcon)
	middlewares := middleware.NewMiddlewares(cnf)

usrSvc := user.NewService(userRepo)
prdSvc := user.NewService(productRepo)

	productHandler := prdHandler.NewHandler(middlewares, prdSvc)
	userRepo := repo.NewUserRepo(dbcon)
	userHandler := usrHandler.NewHandler(middlewares, usrSvc, cnf)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start(*cnf)
0
}
