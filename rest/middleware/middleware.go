package middleware

import "back-end/config"

type Middlewares struct {
	cnd *config.Config
}

func NewMiddlewares(cnf *config.Config) *Middlewares {
	return &Middlewares{
		cnd: cnf,
	}
}
