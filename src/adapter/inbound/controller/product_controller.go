package controller

import (
	"encoding/json"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"fmt"
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

	var product2 domain.Product

	err := json.NewDecoder(r.Body).Decode(&product2)

	fmt.Println(product2)

	product, err := p.productService.Create(product2)

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
