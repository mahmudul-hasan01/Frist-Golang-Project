package product

import (
	"back-end/database"
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

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Bad request", 400)
		return
	}

	newProduct.ID = id

	database.Update(newProduct)

	util.SendData(w, "Product updated successfully", 200)

}
