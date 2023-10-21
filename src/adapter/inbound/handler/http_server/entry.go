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
	productController := controller.NewProductController(e.productService, e.logger)
	orderController := controller.NewOrderController(e.orderService, e.logger)

	router.HandleFunc("/client", Chain(func(w http.ResponseWriter, r *http.Request) { clientController.CreateClient(w, r) }, Method("POST"), Logging()))
	router.HandleFunc("/order", Chain(func(w http.ResponseWriter, r *http.Request) { orderController.CreateOrder(w, r) }, Method("POST"), Logging()))
	router.HandleFunc("/product", Chain(productController.CreateProduct, Method("POST"), Logging()))

	return http.ListenAndServe(":8080", router)
}
