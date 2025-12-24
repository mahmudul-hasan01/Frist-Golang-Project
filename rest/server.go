package rest

import (
	"back-end/config"
	"back-end/rest/handler/product"
	"back-end/rest/handler/user"
	"back-end/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	ProductHandler *product.Handler
	UserHandler    *user.Handler
}

func NewServer(cnf *config.Config, ProductHandler *product.Handler, UserHandler *user.Handler) *Server {
	return &Server{
		cnf:            cnf,
		ProductHandler: ProductHandler,
		UserHandler:    UserHandler,
	}
}

func (server *Server) Start(cnf config.Config) {
	manager := middleware.NewManager()

	// manager.Use(middleware.Logger, middleware.Hudai)

	mux := http.NewServeMux() // router
	// Routes(mux, manager)

	// globalRouter := middleware.PreFlightHandler(mux)

	manager.Use(middleware.Cors, middleware.PreFlight, middleware.Logger)

	wrappedMux := manager.WrapMux(mux)

	server.ProductHandler.RegisterRoutes(mux, manager)
	server.UserHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)

	fmt.Println("Server running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}

	// manager.Use(middleware.Logger, middleware.Hudai)

	// router

	// globalRouter := middleware.PreFlightHandler(mux)

	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
