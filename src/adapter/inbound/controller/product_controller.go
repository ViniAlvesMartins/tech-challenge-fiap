package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/dto"
	"github.com/gorilla/mux"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type ProductController struct {
	productService port.ProductService
	logger         *slog.Logger
}

func NewProductController(productService port.ProductService, logger *slog.Logger) *ProductController {
	return &ProductController{
		productService: productService,
		logger:         logger,
	}
}

func (p *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var productDto dto.ProductDto

	err := json.NewDecoder(r.Body).Decode(&productDto)

	if err != nil {
		p.logger.Error("Unable to decode the request body.  %v", err)
	}

	errValidate := dto.ValidateProduct(productDto)

	fmt.Println(errValidate)

	if len(errValidate.Errors) > 0 {
		p.logger.Error("validate error.  %v", errValidate)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errValidate)
		return
	}

	productDomain := dto.ConvertDtoToDomain(productDto)

	product, err := p.productService.Create(productDomain)

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (p *ProductController) GetProductByCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := mux.Vars(r)["categoryid"]

	categoryIdInt, err := strconv.Atoi(categoryId)

	prod, err := p.productService.GetProductByCategory(categoryIdInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(prod)
	if err != nil {
		return
	}

	

}