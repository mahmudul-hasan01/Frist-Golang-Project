package user

import (
	"back-end/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /user", manager.With(http.HandlerFunc(h.CreateUser), h.middlewares.AuthenticateJwt))

	mux.HandleFunc("GET /user", http.HandlerFunc(h.CreateUser))

}
