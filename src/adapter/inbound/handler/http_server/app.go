package http_server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/controller"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"github.com/gorilla/mux"
)

type App struct {
	logger          *slog.Logger
	clientService   port.ClientService
	productService  port.ProductService
	orderService    port.OrderService
	paymentService  port.PaymentService
	categoryService port.CategoryService
	checkoutService port.CheckoutService
}

func NewApp(logger *slog.Logger,
	clientService port.ClientService,
	productService port.ProductService,
	orderService port.OrderService,
	paymentService port.PaymentService,
	categoryService port.CategoryService,
	checkoutService port.CheckoutService,
) *App {
	return &App{
		logger:          logger,
		clientService:   clientService,
		productService:  productService,
		orderService:    orderService,
		paymentService:  paymentService,
		categoryService: categoryService,
		checkoutService: checkoutService,
	}
}

func (e *App) Run(ctx context.Context) error {
	router := mux.NewRouter()

	clientController := controller.NewClientController(e.clientService, e.logger)
	router.HandleFunc("/client", clientController.CreateClient).Methods("POST")
	router.HandleFunc("/client", clientController.GetClientByCpf).Methods("GET")

	productController := controller.NewProductController(e.productService, e.categoryService, e.logger)
	router.HandleFunc("/product", productController.CreateProduct).Methods("POST")
	router.HandleFunc("/product/{categoryid:[0-9]+}", productController.GetProductByCategory).Methods("GET")
	router.HandleFunc("/product", productController.UpdateProduct).Methods("PATCH")
	router.HandleFunc("/product/{productId:[0-9]+}", productController.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/product/{categoryid:[0-9]+}", productController.GetProductByCategory).Methods("GET")

	orderController := controller.NewOrderController(e.orderService, e.productService, e.logger)
	router.HandleFunc("/order", orderController.FindOrders).Methods("GET")
	router.HandleFunc("/order/{orderId:[0-9]+}", orderController.GetOrderById).Methods("GET")
	router.HandleFunc("/order", orderController.CreateOrder).Methods("POST")

	paymentController := controller.NewPaymentController(e.checkoutService, e.logger)
	router.HandleFunc("/payments", paymentController.CreatePayment).Methods("POST")

	return http.ListenAndServe(":8080", router)
}
