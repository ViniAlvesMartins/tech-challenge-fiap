package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/dto"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/enum"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"log/slog"
	"net/http"
)

type PaymentController struct {
	paymentService port.PaymentService
	orderService   port.OrderService
	logger         *slog.Logger
}

type Response struct {
	MessageError string
}

func NewPaymentController(p port.PaymentService, orderService port.OrderService, logger *slog.Logger) *PaymentController {
	return &PaymentController{
		paymentService: p,
		orderService:   orderService,
		logger:         logger,
	}
}

func (p *PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var err error
	var paymentDTO dto.PaymentDto
	var response Response

	err = json.NewDecoder(r.Body).Decode(&paymentDTO)

	if err != nil {
		p.logger.Error("Unable to decode the request body.  %v", err)
	}

	order, errOrder := p.orderService.GetById(paymentDTO.Order)

	if errOrder != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if order.ID == 0 {
		response = Response{
			MessageError: "Order not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	payment := entity.Payment{
		OrderID: paymentDTO.Order,
		Type:    enum.PaymentType(paymentDTO.Type),
		Amount:  order.Amount,
	}

	err = p.paymentService.Create(&payment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
