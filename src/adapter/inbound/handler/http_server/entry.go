package http_server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/controller"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"github.com/gorilla/mux"
)

type Entry struct {
	logger         *slog.Logger
	clientService  port.ClientService
	productService port.ProductService
}

func NewEntry(logger *slog.Logger, clientService port.ClientService, productService port.ProductService) *Entry {
	return &Entry{
		logger:         logger,
		clientService:  clientService,
		productService: productService,
	}
}

func (e *Entry) Run(ctx context.Context) error {
	router := mux.NewRouter()

	clientController := controller.NewClientController(e.clientService, e.logger)
	productController := controller.NewProductController(e.productService, e.logger)

	router.HandleFunc("/client", Chain(func(w http.ResponseWriter, r *http.Request) { clientController.CreateClient(w, r) }, Method("POST"), Logging()))
	router.HandleFunc("/product", Chain(productController.CreateProduct, Method("POST"), Logging()))
	router.HandleFunc("/product", Chain(func(w http.ResponseWriter, r *http.Request) {
		productController.GetProductByCategory(w, r)
	}, Method("GET"), Logging()))

	return http.ListenAndServe(":8080", router)
}
