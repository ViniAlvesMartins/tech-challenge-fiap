package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer"
	dto "github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer/input"

	"log/slog"
	"net/http"
)

type PaymentController struct {
	paymentUseCase contract.PaymentUseCase
	logger         *slog.Logger
}

func NewPaymentController(p contract.PaymentUseCase, logger *slog.Logger) *PaymentController {
	return &PaymentController{
		paymentUseCase: p,
		logger:         logger,
	}
}

func (p *PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var err error
	var paymentDTO dto.PaymentDto

	err = json.NewDecoder(r.Body).Decode(&paymentDTO)

	if err != nil {
		p.logger.Error("Unable to decode the request body.  %v", slog.Any("error", err))
	}

	errValidate := serializer.Validate(paymentDTO)

	if len(errValidate.Errors) > 0 {
		p.logger.Error("validate error", slog.Any("error", errValidate))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errValidate)
		return
	}

	if err = p.paymentUseCase.Checkout(paymentDTO.Order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (p *PaymentController) GetLastPaymentStatus(w http.ResponseWriter, r *http.Request) {
	var response Response
	orderId := mux.Vars(r)["orderId"]
	orderIdInt, err := strconv.Atoi(orderId)

	if err != nil {
		http.Error(w, "Error to convert id order to int.  %v", http.StatusInternalServerError)
	}

	paymentStatus, err := p.paymentUseCase.GetLastPaymentStatus(orderIdInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response = Response{
		MessageError: "",
		Data:         paymentStatus,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}

}
