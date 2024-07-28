package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/controller/serializer/input"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/controller/serializer/output"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract"
	"github.com/gorilla/mux"
)

type OrderController struct {
	orderUseCase  contract.OrderUseCase
	clientUseCase contract.ClientUseCase
	logger        *slog.Logger
}

func NewOrderController(orderUseCase contract.OrderUseCase, clientUseCase contract.ClientUseCase, logger *slog.Logger) *OrderController {
	return &OrderController{
		orderUseCase:  orderUseCase,
		clientUseCase: clientUseCase,
		logger:        logger,
	}
}

// CreateOrder godoc
// @Summary      Create order
// @Description  Place a new order
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        request   body      input.OrderDto  true  "Order properties"
// @Success      201  {object}  Response{data=output.OrderDto}
// @Failure      500  {object}  swagger.InternalServerErrorResponse{data=interface{}}
// @Failure      404  {object}  swagger.ResourceNotFoundResponse{data=interface{}}
// @Router       /orders [post]
func (o *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderDto input.OrderDto

	if err := json.NewDecoder(r.Body).Decode(&orderDto); err != nil {
		o.logger.Error("unable to decode the request body", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "Unable to decode the request body",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	if len(orderDto.Products) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		jsonResponse, _ := json.Marshal(Response{
			Error: "Product is required",
			Data:  nil,
		})
		w.Write(jsonResponse)
		return
	}

	client, err := o.clientUseCase.GetById(orderDto.ClientId)
	if client == nil && orderDto.ClientId != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "Client not found",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	order, err := o.orderUseCase.Create(orderDto.ConvertToEntity())
	if err != nil {
		o.logger.Error("error creating order", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "Error creating order",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	orderOutput := output.OrderFromEntity(*order)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	jsonResponse, _ := json.Marshal(
		Response{
			Error: "",
			Data:  orderOutput,
		})
	w.Write(jsonResponse)
	return
}

// FindOrders godoc
// @Summary      List orders
// @Description  List orders by status
// @Tags         Orders
// @Produce      json
// @Success      200  {object}  Response{data=[]output.OrderDto}
// @Failure      500  {object}  swagger.InternalServerErrorResponse{data=interface{}}
// @Router       /orders [get]
func (o *OrderController) FindOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := o.orderUseCase.GetAll()
	if err != nil {
		o.logger.Error("error listing orders", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "Error listing orders",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	ordersOutput := output.OrderListFromEntity(*orders)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, _ := json.Marshal(Response{
		Error: "",
		Data:  ordersOutput,
	})
	w.Write(jsonResponse)
	return
}

// GetOrderById godoc
// @Summary      Find order
// @Description  Find order by id
// @Tags         Orders
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Success      200  {object}  Response{data=output.OrderDto}
// @Failure      500  {object}  swagger.InternalServerErrorResponse{data=interface{}}
// @Failure      404  {object}  swagger.ResourceNotFoundResponse{data=interface{}}
// @Router       /orders/{id} [get]
func (o *OrderController) GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderIdParam := mux.Vars(r)["orderId"]

	id, err := strconv.Atoi(orderIdParam)
	if err != nil {
		o.logger.Error("error to convert id order to int", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "Order id must be an integer",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	order, err := o.orderUseCase.GetById(id)
	if err != nil {
		o.logger.Error("error finding order", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "Error finding order",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	if order == nil {
		w.WriteHeader(http.StatusNotFound)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "Order not found",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	orderOutput := output.OrderFromEntity(*order)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, _ := json.Marshal(Response{
		Error: "",
		Data:  orderOutput,
	})
	w.Write(jsonResponse)
	return
}

// UpdateOrderStatusById godoc
// @Summary      Find order
// @Description  Find order by id
// @Tags         Orders
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Param        request   body      input.StatusOrderDto  true  "Order status"
// @Success      204  {object}  interface{}
// @Failure      500  {object}  swagger.InternalServerErrorResponse{data=interface{}}
// @Failure      404  {object}  swagger.ResourceNotFoundResponse{data=interface{}}
// @Router       /orders/{id} [patch]
func (o *OrderController) UpdateOrderStatusById(w http.ResponseWriter, r *http.Request) {
	orderIdParam := mux.Vars(r)["orderId"]
	orderId, err := strconv.Atoi(orderIdParam)

	if err != nil {
		o.logger.Error("error converting orderId to int", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse, _ := json.Marshal(Response{
			Error: "Error to convert orderId to int",
			Data:  nil,
		})
		w.Write(jsonResponse)
		return
	}

	var statusOrderDto input.StatusOrderDto
	if err := json.NewDecoder(r.Body).Decode(&statusOrderDto); err != nil {
		o.logger.Error("unable to decode the request body", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "Unable to decode the request body",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	if !enum.ValidateOrderStatus(statusOrderDto.Status) {
		w.WriteHeader(http.StatusBadRequest)
		jsonResponse, _ := json.Marshal(Response{
			Error: "Invalid status",
			Data:  nil,
		})
		w.Write(jsonResponse)
		return
	}

	order, err := o.orderUseCase.GetById(orderId)
	if err != nil {
		o.logger.Error("error getting order by id", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse, _ := json.Marshal(Response{
			Error: "Error to get order",
			Data:  nil,
		})
		w.Write(jsonResponse)
		return
	}

	if order == nil {
		w.WriteHeader(http.StatusNotFound)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "Order not found",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	if err := o.orderUseCase.UpdateStatusById(orderId, enum.StatusOrder(statusOrderDto.Status)); err != nil {
		o.logger.Error("error updating status by id", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse, _ := json.Marshal(Response{
			Error: "Error updating status",
			Data:  nil,
		})
		w.Write(jsonResponse)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
