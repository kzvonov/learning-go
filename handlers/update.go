package handlers

import (
	"net/http"

	"github.com/kzvonov/learning-go/data"
)

// swagger:route PUT /products/{id} products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	product := r.Context().Value(KeyProduct{}).(*data.Product)
	err := data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found"}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
