package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer"
	dto "github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer/input"

	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
	productUseCase  contract.ProductUseCase
	categoryUseCase contract.CategoryUseCase
	logger          *slog.Logger
}

func NewProductController(productUseCase contract.ProductUseCase, categoryUseCase contract.CategoryUseCase, logger *slog.Logger) *ProductController {
	return &ProductController{
		productUseCase:  productUseCase,
		logger:          logger,
		categoryUseCase: categoryUseCase,
	}
}

func (p *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.ProductDto

	if err := json.NewDecoder(r.Body).Decode(&productDto); err != nil {
		p.logger.Error("Unable to decode the request body.  %v", slog.Any("error", err))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error decoding request body",
				Data:         nil,
			})
		return
	}

	if serialize := serializer.Validate(productDto); len(serialize.Errors) > 0 {
		p.logger.Error("validate error", slog.Any("error", serialize))

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Make all required fields are sent correctly",
				Data:         nil,
			})
		return
	}

	category, err := p.categoryUseCase.GetById(productDto.CategoryId)
	if err != nil {
		p.logger.Error("validate getting category by id", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error getting category",
				Data:         nil,
			})
		return
	}

	if category == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Category not found",
				Data:         nil,
			})
		return
	}

	product, err := p.productUseCase.Create(productDto.ConvertToEntity())
	if err != nil {
		p.logger.Error("error creating product", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error creating product",
				Data:         nil,
			})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		Response{
			ErrorMessage: "",
			Data:         product,
		})
}

func (p *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.ProductDto

	productIdParam, ok := mux.Vars(r)["productId"]
	if !ok {
		p.logger.Error("id is missing in parameters")

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Id is missing in parameters",
				Data:         nil,
			})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&productDto); err != nil {
		p.logger.Error("Unable to decode the request body.  %v", slog.Any("error", err))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error decoding request body",
				Data:         nil,
			})
		return
	}

	productId, err := strconv.Atoi(productIdParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Id is not a number",
				Data:         nil,
			})
		return
	}

	if serialize := serializer.Validate(productDto); len(serialize.Errors) > 0 {
		p.logger.Error("validate error", slog.Any("error", serialize))

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Make sure all required fields are sent correctly",
				Data:         nil,
			})
		return
	}

	validateProduct, err := p.productUseCase.GetById(productId)
	if err != nil {
		p.logger.Error("error getting product by id", slog.Any("error", err))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error finding product",
				Data:         nil,
			})
		return
	}

	if validateProduct == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Product not found",
				Data:         nil,
			})
		return
	}

	productDto.SetID(productId)
	productDomain := productDto.ConvertToEntity()

	product, err := p.productUseCase.Update(productDomain)
	if err != nil {
		p.logger.Error("error updating product data", slog.Any("error", err))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error updating product data",
				Data:         nil,
			})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		Response{
			ErrorMessage: "",
			Data:         product,
		})
}

func (p *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productIdParam, ok := mux.Vars(r)["productId"]
	if !ok {
		p.logger.Error("id is missing in parameters")

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Id is missing in parameters",
				Data:         nil,
			})
		return
	}

	productId, err := strconv.Atoi(productIdParam)
	if err != nil {
		p.logger.Error("Error to convert productId to int.  %v", err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Id is not a number",
				Data:         nil,
			})
		return
	}

	validateProduct, err := p.productUseCase.GetById(productId)
	if err != nil {
		p.logger.Error("error getting product by id", slog.Any("error", err))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error finding product",
				Data:         nil,
			})
		return
	}

	if validateProduct == nil || validateProduct.Active == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Product not found",
				Data:         nil,
			})
		return
	}

	if err := p.productUseCase.Delete(productId); err != nil {
		p.logger.Error("error deleting product", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error deleting product",
				Data:         nil,
			})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (p *ProductController) GetProductByCategory(w http.ResponseWriter, r *http.Request) {
	categoryIdParam := mux.Vars(r)["categoryId"]

	categoryId, err := strconv.Atoi(categoryIdParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Id is missing in parameters",
				Data:         nil,
			})
		return
	}

	category, err := p.categoryUseCase.GetById(categoryId)
	if err != nil {
		p.logger.Error("error getting category by id", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error finding category",
				Data:         nil,
			})
		return
	}

	if category == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Category not found",
				Data:         nil,
			})
		return
	}

	products, err := p.productUseCase.GetProductByCategory(categoryId)
	if err != nil {
		p.logger.Error("error getting products by category", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error finding products",
				Data:         nil,
			})
		return
	}

	if len(products) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Product not found",
				Data:         nil,
			})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		Response{
			ErrorMessage: "",
			Data:         products,
		})
}
