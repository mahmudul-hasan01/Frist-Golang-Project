package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	GlobalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		GlobalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) *Manager {
	mngr.GlobalMiddlewares = append(mngr.GlobalMiddlewares, middlewares...)
	return mngr
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, globalMiddleware := range mngr.GlobalMiddlewares {
		n = globalMiddleware(n)
	}

	return n
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler

	for _, middleware := range mngr.GlobalMiddlewares {
		h = middleware(h)
	}

	// for _, globalMiddleware := range mngr.GlobalMiddlewares {
	// 	n = globalMiddleware(n)
	// }

	return h
}
