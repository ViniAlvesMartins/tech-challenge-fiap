package controller

import (
	"encoding/json"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer/input"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer/output"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"

	"github.com/gorilla/mux"
)

type OrderController struct {
	orderUseCase   contract.OrderUseCase
	productUseCase contract.ProductUseCase
	clientUseCase  contract.ClientUseCase
	logger         *slog.Logger
}

func NewOrderController(orderUseCase contract.OrderUseCase, productUseCase contract.ProductUseCase, clientUseCase contract.ClientUseCase, logger *slog.Logger) *OrderController {
	return &OrderController{
		orderUseCase:   orderUseCase,
		productUseCase: productUseCase,
		clientUseCase:  clientUseCase,
		logger:         logger,
	}
}

func (o *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderDto input.OrderDto

	if err := json.NewDecoder(r.Body).Decode(&orderDto); err != nil {
		o.logger.Error("unable to decode the request body", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Unable to decode the request body",
				Data:  nil,
			})
		return
	}

	if prods := len(orderDto.Products); prods < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Error: "Product is required",
			Data:  nil,
		})
		return
	}

	var products []*entity.Product
	for _, p := range orderDto.Products {
		product, err := o.productUseCase.GetById(p.ID)

		if err != nil {
			o.logger.Error("error getting product by id", slog.String("message", err.Error()))

			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{
				Error: "Error finding product",
				Data:  nil,
			})
			return
		}

		if product == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(
				Response{
					Error: fmt.Sprintf("Product of id %d not found", p.ID),
					Data:  nil,
				})
			return
		}

		products = append(products, product)
	}

	client, err := o.clientUseCase.GetClientById(orderDomain.ClientId)

	if client == nil && orderDomain.ClientId != nil {
		response = Response{
			MessageError: "Client not exists",
			Data:         nil,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	order, err := o.orderUseCase.Create(orderDto.ConvertToEntity(), products)
	if err != nil {
		o.logger.Error("error creating order", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Error creating order",
				Data:  nil,
			})
		return
	}

	orderOutput := output.OrderFromEntity(*order)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		Response{
			Error: "",
			Data:  orderOutput,
		})
}

func (o *OrderController) FindOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := o.orderUseCase.GetAll()
	if err != nil {
		o.logger.Error("error listing orders", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Error listing orders",
				Data:  nil,
			})
		return
	}

	ordersOutput := output.OrderListFromEntity(*orders)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Error: "",
		Data:  ordersOutput,
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
				Error: "Order id must be an integer",
				Data:  nil,
			})
		return
	}

	order, err := o.orderUseCase.GetById(id)
	if err != nil {
		o.logger.Error("error finding order", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Error finding order",
				Data:  nil,
			})
		return
	}

	if order == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Order not found",
				Data:  nil,
			})
		return
	}

	orderOutput := output.OrderFromEntity(*order)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Error: "",
		Data:  orderOutput,
	})
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func (o *OrderController) UpdateOrderStatusById(w http.ResponseWriter, r *http.Request) {
	var response Response
	orderId := mux.Vars(r)["orderId"]
	orderIdInt, err := strconv.Atoi(orderId)

	if err != nil {
		http.Error(w, "Error to convert id order to int.  %v", http.StatusInternalServerError)
	}

	status := r.URL.Query().Get("status")
	validateStatus, err := enum.ValidateStatus(status)

	if validateStatus == false {
		response = Response{
			MessageError: err.Error(),
			Data:         nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	order, errOrder := o.orderUseCase.GetById(orderIdInt)

	if errOrder != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if order == nil {
		response = Response{
			MessageError: "Not found",
			Data:         nil,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = o.orderUseCase.UpdateStatusById(orderIdInt, enum.StatusOrder(status))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
