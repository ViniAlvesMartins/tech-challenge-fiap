package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
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
	var response Response

	err := json.NewDecoder(r.Body).Decode(&orderDomain)

	if err != nil {
		o.logger.Error("Unable to decode the request body.  %v", err)
	}

	validateLengthProds := len(orderDomain.Products)

	if validateLengthProds <= 0 {
		response = Response{
			MessageError: "Product is required",
			Data:         nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	var products []*entity.Product
	for _, product := range orderDomain.Products {
		prod, errProd := o.productUseCase.GetById(product.ID)

		if errProd != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if prod == nil {
			response = Response{
				MessageError: "Product not found " + strconv.Itoa(product.ID),
				Data:         nil,
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}

		products = append(products, prod)
	}

	order, err := o.orderUseCase.Create(orderDomain, products)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response = Response{
		MessageError: "",
		Data:         order,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func (o *OrderController) FindOrders(w http.ResponseWriter, r *http.Request) {
	var orders *[]entity.Order

	orders, err := o.orderUseCase.GetAll()

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
	var response Response
	orderId := mux.Vars(r)["orderId"]
	orderIdInt, err := strconv.Atoi(orderId)

	if err != nil {
		http.Error(w, "Error to convert id order to int.  %v", http.StatusInternalServerError)
	}

	order, err := o.orderUseCase.GetById(orderIdInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if order == nil {
		response = Response{
			MessageError: "Order Not found",
			Data:         nil,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response = Response{
		MessageError: "",
		Data:         order,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
