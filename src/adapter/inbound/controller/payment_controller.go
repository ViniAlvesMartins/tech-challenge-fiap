package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/dto"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"log/slog"
	"net/http"
)

type PaymentController struct {
	checkoutService port.CheckoutService
	logger          *slog.Logger
}

func NewPaymentController(c port.CheckoutService, logger *slog.Logger) *PaymentController {
	return &PaymentController{
		checkoutService: c,
		logger:          logger,
	}
}

func (p *PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var err error
	var paymentDTO dto.PaymentDto

	err = json.NewDecoder(r.Body).Decode(&paymentDTO)

	if err != nil {
		p.logger.Error("Unable to decode the request body.  %v", err)
	}

	errValidate := dto.Validate(paymentDTO)

	if len(errValidate.Errors) > 0 {
		p.logger.Error("validate error.  %v", errValidate)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errValidate)
		return
	}

	if err = p.checkoutService.PayWithQRCode(paymentDTO.Order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
