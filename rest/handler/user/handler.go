package user

import (
	"back-end/config"
	"back-end/repo"
	"back-end/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	userRepo    repo.UserRepo
	cnf         *config.Config
}

func NewHandler(middlewares *middleware.Middlewares, userRepo repo.UserRepo, cnf *config.Config) *Handler {
	return &Handler{
		middlewares: middlewares,
		userRepo:    userRepo,
		cnf:         cnf,
	}
}
