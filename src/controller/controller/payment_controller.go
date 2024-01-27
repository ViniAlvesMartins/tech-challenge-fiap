package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer"
	dto "github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer/input"
	"log/slog"
	"net/http"
)

type PaymentController struct {
	paymentUseCase contract.PaymentUseCase
	logger         *slog.Logger
	orderUseCase   contract.OrderUseCase
}

type GetLastPaymentStatus struct {
	OrderId       int
	PaymentStatus enum.PaymentStatus
}

func NewPaymentController(p contract.PaymentUseCase, logger *slog.Logger, orderUseCase contract.OrderUseCase) *PaymentController {
	return &PaymentController{
		paymentUseCase: p,
		logger:         logger,
		orderUseCase:   orderUseCase,
	}
}

func (p *PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var paymentDTO dto.PaymentDto
	var response Response

	orderIdParam := mux.Vars(r)["orderId"]
	orderId, err := strconv.Atoi(orderIdParam)
	if err != nil {
		p.logger.Error("error to convert id order to int", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Order id must be an integer",
				Data:  nil,
			})
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&paymentDTO); err != nil {
		p.logger.Error("Unable to decode the request body.  %v", slog.Any("error", err))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Error decoding request body",
				Data:  nil,
			})
		return
	}

	if serialize := serializer.Validate(paymentDTO); len(serialize.Errors) > 0 {
		p.logger.Error("validate error", slog.Any("error", serialize))

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Make sure all required fields are sent correctly",
				Data:  nil,
			})
		return
	}

	order, err := p.orderUseCase.GetById(orderId)
	if err != nil {
		p.logger.Error("error getting order", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Error getting order details",
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

	qrCode, err := p.paymentUseCase.CreateQRCode(order)
	if err != nil {
		p.logger.Error("error creating qr code", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Error creating qr code",
				Data:  nil,
			})

		return
	}

	if qrCode == nil {
		response = Response{
			Error: "O pagamento para o pedido já foi efetuado",
			Data:  "",
		}
	} else {
		response = Response{
			Error: "",
			Data:  qrCode,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		Response{
			Error: "",
			Data:  response,
		})
}

func (p *PaymentController) GetLastPaymentStatus(w http.ResponseWriter, r *http.Request) {
	orderIdParam := mux.Vars(r)["orderId"]
	orderId, err := strconv.Atoi(orderIdParam)
	if err != nil {
		p.logger.Error("error to convert id order to int", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Order id must be an integer",
				Data:  nil,
			})
		return
	}

	paymentStatus, err := p.paymentUseCase.GetLastPaymentStatus(orderId)
	if err != nil {
		p.logger.Error("error getting last payment status", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Error getting last payment status",
				Data:  nil,
			})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		Response{
			Error: "",
			Data: GetLastPaymentStatus{
				OrderId:       orderId,
				PaymentStatus: paymentStatus,
			},
		})
}

func (p *PaymentController) Notification(w http.ResponseWriter, r *http.Request) {
	orderIdParam := mux.Vars(r)["orderId"]

	orderId, err := strconv.Atoi(orderIdParam)
	if err != nil {
		p.logger.Error("error to convert id order to int", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Order id must be an integer",
				Data:  nil,
			})
		return
	}

	order, err := p.orderUseCase.GetById(orderId)
	if err != nil {
		p.logger.Error("error getting order by id", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Error getting order by id",
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

	if err = p.paymentUseCase.PaymentNotification(order); err != nil {
		p.logger.Error("error processing payment notification", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			Response{
				Error: "Error processing payment notification",
				Data:  nil,
			})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
