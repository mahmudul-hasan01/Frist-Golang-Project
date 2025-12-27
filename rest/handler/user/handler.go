package user

import (
	"back-end/config"
	"back-end/repo"
	"back-end/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	svc Service
	cnf         *config.Config
}

func NewHandler(middlewares *middleware.Middlewares, svc Service, cnf *config.Config) *Handler {
	return &Handler{
		middlewares: middlewares,
		svc: svc
		cnf:         cnf,

	}
}
