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
	router.HandleFunc("/client", clientController.CreateClient).Methods("POST")
	router.HandleFunc("/client", clientController.GetClientByCpf).Methods("GET")

	productController := controller.NewProductController(e.productService, e.logger)
	router.HandleFunc("/product", productController.CreateProduct).Methods("POST")
  router.HandleFunc("/product/{categoryid:[0-9]+}", productController.GetProductByCategory).Methods("GET")

	orderController := controller.NewOrderController(e.orderService, e.logger)
	router.HandleFunc("/order", orderController.FindOrders).Methods("GET")
	router.HandleFunc("/order", orderController.CreateOrder).Methods("POST")

	return http.ListenAndServe(":8080", router)
}