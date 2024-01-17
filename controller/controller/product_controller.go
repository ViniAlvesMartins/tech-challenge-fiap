package controller

import (
	"encoding/json"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/controller/serializer"
	"log/slog"
	"net/http"
	"strconv"

	dto "github.com/ViniAlvesMartins/tech-challenge-fiap/controller/serializer/input"
	"github.com/gorilla/mux"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type ProductController struct {
	productService  port.ProductService
	categoryService port.CategoryService
	logger          *slog.Logger
}

func NewProductController(productService port.ProductService, categoryService port.CategoryService, logger *slog.Logger) *ProductController {
	return &ProductController{
		productService:  productService,
		logger:          logger,
		categoryService: categoryService,
	}
}

func (p *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.ProductDto
	var response Response

	err := json.NewDecoder(r.Body).Decode(&productDto)

	if err != nil {
		p.logger.Error("Unable to decode the request body.  %v", err)
	}

	errValidate := serializer.Validate(productDto)

	fmt.Println(errValidate)

	if len(errValidate.Errors) > 0 {
		p.logger.Error("validate error", slog.Any("error", errValidate))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errValidate)
		return
	}

	category, errCategory := p.categoryService.GetById(productDto.CategoryId)

	if errCategory != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if category == nil {
		response = Response{
			MessageError: "Category not found",
			Data:         nil,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	p.logger.Info("category")
	fmt.Println(category != nil)

	productDomain := productDto.ConvertToEntity()

	product, err := p.productService.Create(productDomain)

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response = Response{
		MessageError: "",
		Data:         product,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (p *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var response Response
	vars := mux.Vars(r)

	productIdParam, ok := vars["productId"]

	if !ok {
		response := Response{
			MessageError: "id is missing in parameters",
			Data:         nil,
		}

		p.logger.Error("id is missing in parameters")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	var productDto dto.ProductDto

	err := json.NewDecoder(r.Body).Decode(&productDto)

	if err != nil {
		p.logger.Error("Unable to decode the request body.  %v", err)
	}

	convertId, err := strconv.Atoi(productIdParam)

	if err != nil {
		response := Response{
			MessageError: "Id not is number",
			Data:         nil,
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	errValidate := serializer.Validate(productDto)

	fmt.Println(errValidate)

	if len(errValidate.Errors) > 0 {
		p.logger.Error("validate error", slog.Any("error", errValidate))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errValidate)
		return
	}

	validateProduct, errProduct := p.productService.GetById(convertId)

	if errProduct != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if validateProduct == nil {
		response = Response{
			MessageError: "Not found",
			Data:         nil,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	productDto.SetID(convertId)
	productDomain := productDto.ConvertToEntity()
	product, err := p.productService.Update(productDomain)

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response = Response{
		MessageError: "",
		Data:         product,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (p *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var response Response
	vars := mux.Vars(r)
	productIdParam, ok := vars["productId"]

	if !ok {
		fmt.Println("id is missing in parameters")
	}

	productId, err := strconv.Atoi(productIdParam)

	if err != nil {
		p.logger.Error("Error to convert productId to int.  %v", err)
	}

	validateProduct, errProduct := p.productService.GetById(productId)

	if errProduct != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if validateProduct == nil || validateProduct.Active == false {
		response = Response{
			MessageError: "Not found",
			Data:         nil,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = p.productService.Delete(productId)
	if err != nil {
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		return
	}
}

func (p *ProductController) GetProductByCategory(w http.ResponseWriter, r *http.Request) {
	var response Response
	categoryId := mux.Vars(r)["categoryId"]

	categoryIdInt, ok := strconv.Atoi(categoryId)

	if ok != nil {
		response := Response{
			MessageError: "id is missing in parameters",
			Data:         nil,
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	category, errCateg := p.categoryService.GetById(categoryIdInt)

	if errCateg != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if category == nil {
		response = Response{
			MessageError: "Category Not found",
			Data:         nil,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	prod, err := p.productService.GetProductByCategory(categoryIdInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	lenProd := len(prod)

	if lenProd == 0 {
		response = Response{
			MessageError: "Product Not found",
			Data:         nil,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response = Response{
		MessageError: "",
		Data:         prod,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
