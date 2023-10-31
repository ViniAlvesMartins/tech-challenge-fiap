package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type OrderController struct {
	orderService   port.OrderService
	productService port.ProductService
	logger         *slog.Logger
}

func NewOrderController(orderService port.OrderService, productService port.ProductService, logger *slog.Logger) *OrderController {
	return &OrderController{
		orderService:   orderService,
		productService: productService,
		logger:         logger,
	}
}

func (o *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderDomain entity.Order
	var response Response

	err := json.NewDecoder(r.Body).Decode(&orderDomain)

	if err != nil {
		o.logger.Error("Unable to decode the request body.  %v", err)
	}

	for _, product := range orderDomain.Products {

		prod, errProd := o.productService.GetById(product.ID)

		if errProd != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if prod.ID == 0 {
			response = Response{
				MessageError: "Product not found " + strconv.Itoa(product.ID),
				Data:         nil,
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}

	}

	order, err := o.orderService.Create(orderDomain)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		return
	}
}

func (o *OrderController) FindOrders(w http.ResponseWriter, r *http.Request) {
	var orders *[]entity.Order

	orders, err := o.orderService.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		return
	}
}

func (o *OrderController) GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderId := mux.Vars(r)["orderId"]
	orderIdInt, err := strconv.Atoi(orderId)

	if err != nil {
		http.Error(w, "Error to convert id order to int.  %v", http.StatusInternalServerError)
	}

	order, err := o.orderService.GetById(orderIdInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if order == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		return
	}
}
