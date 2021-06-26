package handlers

import (
	"net/http"

	"github.com/kzvonov/learning-go/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns: a list of products
// responses:
//   201: noContent

// DeleteProduct deletes a product
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
