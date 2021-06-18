package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"../data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		idRegExp := regexp.MustCompile(`/([0-9]+)`)

		g := idRegExp.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) < 1 {
			p.l.Println("more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("cant conver str to int")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.l.Println("got id:", id)
		p.updateProduct(id, rw, r)
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	product := &data.Product{}

	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode JSON", http.StatusBadRequest)
	}

	data.AddProduct(product)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	product := &data.Product{}

	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode JSON", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}
}
