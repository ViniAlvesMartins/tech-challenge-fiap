package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/gorilla/mux"
)

type OrderController struct {
	orderUseCase   contract.OrderUseCase
	productUseCase contract.ProductUseCase
	logger         *slog.Logger
}

func NewOrderController(orderUseCase contract.OrderUseCase, productUseCase contract.ProductUseCase, logger *slog.Logger) *OrderController {
	return &OrderController{
		orderUseCase:   orderUseCase,
		productUseCase: productUseCase,
		logger:         logger,
	}
}

func (o *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderDomain entity.Order

	if err := json.NewDecoder(r.Body).Decode(&orderDomain); err != nil {
		o.logger.Error("unable to decode the request body", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Unable to decode the request body",
				Data:         nil,
			})
		return
	}

	if prods := len(orderDomain.Products); prods < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			ErrorMessage: "Product is required",
			Data:         nil,
		})
		return
	}

	var products []*entity.Product
	for _, p := range orderDomain.Products {
		product, err := o.productUseCase.GetById(p.ID)

		if err != nil {
			o.logger.Error("error getting product by id", slog.String("message", err.Error()))

			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{
				ErrorMessage: "Error finding product",
				Data:         nil,
			})
			return
		}

		if product == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(
				Response{
					ErrorMessage: fmt.Sprintf("Product of id %d not found", p.ID),
					Data:         nil,
				})
			return
		}

		products = append(products, product)
	}

	order, err := o.orderUseCase.Create(orderDomain, products)
	if err != nil {
		o.logger.Error("error creating order", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error creating order",
				Data:         nil,
			})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		Response{
			ErrorMessage: "",
			Data:         order,
		})
}

func (o *OrderController) FindOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := o.orderUseCase.GetAll()
	if err != nil {
		o.logger.Error("error listing orders", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error listing orders",
				Data:         nil,
			})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		ErrorMessage: "",
		Data:         orders,
	})
}

func (o *OrderController) GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderIdParam := mux.Vars(r)["orderId"]

	id, err := strconv.Atoi(orderIdParam)
	if err != nil {
		o.logger.Error("error to convert id order to int", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Order id must be an integer",
				Data:         nil,
			})
		return
	}

	order, err := o.orderUseCase.GetById(id)
	if err != nil {
		o.logger.Error("error finding order", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Error finding order",
				Data:         nil,
			})
		return
	}

	if order == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			Response{
				ErrorMessage: "Order not found",
				Data:         nil,
			})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		ErrorMessage: "",
		Data:         order,
	})
}
