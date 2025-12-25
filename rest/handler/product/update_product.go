package product

import (
	"back-end/repo"
	"back-end/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProducts(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Check if product exists
	existingProduct, err := h.productRepo.Get(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if existingProduct == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	var updatedProduct repo.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	updatedProduct.ID = id

	result, err := h.productRepo.Update(updatedProduct)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, result, http.StatusOK)
}
