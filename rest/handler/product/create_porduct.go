package product

import (
	"back-end/repo"
	"back-end/util"
	"encoding/json"
	"net/http"
)

type ReqCreateProduct struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImgUrl      string  `json:"imageUrl"`
	Price       float64 `json:"price"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct ReqCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Convert ReqCreateProduct to repo.Product
	product := domain.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		ImgUrl:      newProduct.ImgUrl,
		Price:       newProduct.Price,
	}

	createdProduct, err := h.svc.Create(product)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdProduct, http.StatusCreated)
}
