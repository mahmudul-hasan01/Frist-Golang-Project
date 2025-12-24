package product

import (
	"back-end/database"
	"back-end/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")

	id, errr := strconv.Atoi(productId)

	if errr != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := database.Get(id)

	if product == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}
	util.SendData(w, "Product not found", http.StatusNotFound)
}
