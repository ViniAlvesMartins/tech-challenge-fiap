package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type OrderController struct {
	orderService port.OrderService
	logger       *slog.Logger
}

func NewOrderController(orderService port.OrderService, logger *slog.Logger) *OrderController {
	return &OrderController{
		orderService: orderService,
		logger:       logger,
	}
}

func (o *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderDomain entity.Order

	err := json.NewDecoder(r.Body).Decode(&orderDomain)

	if err != nil {
		o.logger.Error("Unable to decode the request body.  %v", err)
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
