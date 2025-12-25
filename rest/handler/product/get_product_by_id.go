package product

import (
	"back-end/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.Get(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if product == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	util.SendData(w, product, http.StatusOK)
}
