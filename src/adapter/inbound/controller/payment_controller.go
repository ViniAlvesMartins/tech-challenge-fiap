package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/dto"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"log/slog"
	"net/http"
)

type PaymentController struct {
	paymentService port.PaymentService
	logger         *slog.Logger
}

func NewPaymentController(p port.PaymentService, logger *slog.Logger) *PaymentController {
	return &PaymentController{
		paymentService: p,
		logger:         logger,
	}
}

func (p *PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var err error
	var paymentDTO dto.PaymentDto

	err = json.NewDecoder(r.Body).Decode(&paymentDTO)

	if err != nil {
		p.logger.Error("Unable to decode the request body.  %v", err)
	}

	err = p.paymentService.Create(paymentDTO.ConvertToEntity())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
