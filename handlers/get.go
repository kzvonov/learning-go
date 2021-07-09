package handlers

import (
	"net/http"

	"github.com/kzvonov/learning-go/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//		200: productsResponse

// Get handles GET requests and returns all current products
func (p *Products) Get(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	rw.Header().Add("Content-Type", "application/json")

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products getProduct
// Return a list of products from the database
// responses:
//		200: productResponse
// 		404: errorResponse

// GetOne handles get request and returns one product
func (p *Products) GetOne(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	product, err := data.GetProductByID(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(product, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
