package controller

import (
	"encoding/json"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"net/http"
)

type ProductController struct {
	productService port.ProductService
}

func NewProductController(productService port.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (p *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {

	product, err := p.productService.Create("teste", 15.15, "bebida")

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
