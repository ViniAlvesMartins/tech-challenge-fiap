package controller

import (
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer"
	dto "github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer/input"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"

	"log/slog"
	"net/http"
)

type PaymentController struct {
	paymentUseCase contract.PaymentUseCase
	logger         *slog.Logger
	orderUseCase   contract.OrderUseCase
}

func NewPaymentController(p contract.PaymentUseCase, logger *slog.Logger, orderUseCase contract.OrderUseCase) *PaymentController {
	return &PaymentController{
		paymentUseCase: p,
		logger:         logger,
		orderUseCase:   orderUseCase,
	}
}

func (p *PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var err error
	var paymentDTO dto.PaymentDto

	orderId := mux.Vars(r)["orderId"]

	orderIdInt, err := strconv.Atoi(orderId)

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

	order, err := p.orderUseCase.GetById(orderIdInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if order == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	qrCode, error := p.paymentUseCase.CreateQRCode(order)

	if error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		MessageError: "",
		Data:         qrCode,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	if err != nil {
		return
	}
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
		Data: GetLastPaymentStatus{
			OrderId:       orderIdInt,
			PaymentStatus: paymentStatus,
		},
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}

}

type GetLastPaymentStatus struct {
	OrderId       int
	PaymentStatus enum.PaymentStatus
}
