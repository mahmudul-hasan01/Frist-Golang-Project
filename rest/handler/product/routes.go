package product

import (
	"back-end/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts), middleware.PreFlight, middleware.Logger))

	mux.Handle("GET /products/:id", http.HandlerFunc(h.GetProductById))

	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct), h.middlewares.AuthenticateJwt))

	mux.Handle("PUT /products/:id", http.HandlerFunc(h.UpdateProducts))

	mux.HandleFunc("DELETE /products/:id", http.HandlerFunc(h.DeleteProduct))

	// mux.HandleFunc("GET /user", http.HandlerFunc(handler.CreateUser))

}
