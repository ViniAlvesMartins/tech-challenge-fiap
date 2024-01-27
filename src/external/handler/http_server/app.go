package http_server

import (
	"context"
	"log/slog"
	"net/http"

	_ "github.com/ViniAlvesMartins/tech-challenge-fiap/doc/swagger"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/controller"
	"github.com/swaggo/http-swagger/v2"

	"github.com/gorilla/mux"
)

type App struct {
	logger          *slog.Logger
	clientUseCase   contract.ClientUseCase
	productUseCase  contract.ProductUseCase
	orderUseCase    contract.OrderUseCase
	paymentUseCase  contract.PaymentUseCase
	categoryUseCase contract.CategoryUseCase
}

func NewApp(logger *slog.Logger,
	clientUseCase contract.ClientUseCase,
	productUseCase contract.ProductUseCase,
	orderUseCase contract.OrderUseCase,
	paymentUseCase contract.PaymentUseCase,
	categoryUseCase contract.CategoryUseCase,
) *App {
	return &App{
		logger:          logger,
		clientUseCase:   clientUseCase,
		productUseCase:  productUseCase,
		orderUseCase:    orderUseCase,
		paymentUseCase:  paymentUseCase,
		categoryUseCase: categoryUseCase,
	}
}

func (e *App) Run(ctx context.Context) error {
	router := mux.NewRouter()

	clientController := controller.NewClientController(e.clientUseCase, e.logger)
	router.HandleFunc("/clients", clientController.CreateClient).Methods("POST")
	router.HandleFunc("/clients", clientController.GetClientByCpf).Methods("GET")

	productController := controller.NewProductController(e.productUseCase, e.categoryUseCase, e.logger)
	router.HandleFunc("/products", productController.CreateProduct).Methods("POST")
	router.HandleFunc("/categories/{categoryId:[0-9]+}/products", productController.GetProductByCategory).Methods("GET")
	router.HandleFunc("/products/{productId:[0-9]+}", productController.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{productId:[0-9]+}", productController.DeleteProduct).Methods("DELETE")

	orderController := controller.NewOrderController(e.orderUseCase, e.productUseCase, e.clientUseCase, e.logger)
	router.HandleFunc("/orders", orderController.FindOrders).Methods("GET")
	router.HandleFunc("/orders/{orderId:[0-9]+}", orderController.GetOrderById).Methods("GET")
	router.HandleFunc("/orders", orderController.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{orderId:[0-9]+}", orderController.UpdateOrderStatusById).Methods("PATCH")

	paymentController := controller.NewPaymentController(e.paymentUseCase, e.logger, e.orderUseCase)
	router.HandleFunc("/orders/{orderId:[0-9]+}/payments", paymentController.CreatePayment).Methods("POST")
	router.HandleFunc("/orders/{orderId:[0-9]+}/status-payment", paymentController.GetLastPaymentStatus).Methods("GET")
	router.HandleFunc("/orders/{orderId:[0-9]+}/notification-payments", paymentController.Notification).Methods("POST")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return http.ListenAndServe(":8080", router)
}
