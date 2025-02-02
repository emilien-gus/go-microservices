package handler

import (
	"encoding/json"
	"log"
	"net/http"

	data "example.com/go-demo-1/product-api/Data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		pl := data.GetProducts()
		e := json.NewEncoder(rw)
		e.Encode(pl)
	}
}
