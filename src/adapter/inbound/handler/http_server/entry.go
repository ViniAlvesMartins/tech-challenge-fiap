package http_server

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/controller"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type Entry struct {
	logger         *slog.Logger
	clientService  port.ClientService
	productService port.ProductService
	orderService   port.OrderService
}

func NewEntry(logger *slog.Logger,
	clientService port.ClientService,
	productService port.ProductService,
	orderService port.OrderService) *Entry {

	return &Entry{
		logger:         logger,
		clientService:  clientService,
		productService: productService,
		orderService:   orderService,
	}
}

func (e *Entry) Run(ctx context.Context) error {
	router := mux.NewRouter()

	clientController := controller.NewClientController(e.clientService, e.logger)
	router.HandleFunc("/client", Chain(clientController.CreateClient, Method("POST"), Logging()))
	router.HandleFunc("/client", Chain(clientController.GetClientByCpf, Method("GET"), Logging()))

	productController := controller.NewProductController(e.productService, e.logger)
	router.HandleFunc("/product", Chain(productController.CreateProduct, Method("POST"), Logging()))

	orderController := controller.NewOrderController(e.orderService, e.logger)
	router.HandleFunc("/order", Chain(orderController.CreateOrder, Method("POST"), Logging()))
	router.HandleFunc("/orders", Chain(orderController.FindOrders, Method("GET"), Logging()))

	return http.ListenAndServe(":8080", router)
}
