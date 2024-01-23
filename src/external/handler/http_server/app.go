package http_server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/controller"

	"github.com/gorilla/mux"
)

type App struct {
	logger          *slog.Logger
	clientUseCase   contract.ClientUseCase
	productUseCase  contract.ProductUseCase
	orderUseCase    contract.OrderUseCase
	paymentUseCase  contract.PaymentUseCase
	categoryUseCase contract.CategoryUseCase
	checkoutUseCase contract.CheckoutUseCase
}

func NewApp(logger *slog.Logger,
	clientUseCase contract.ClientUseCase,
	productUseCase contract.ProductUseCase,
	orderUseCase contract.OrderUseCase,
	paymentUseCase contract.PaymentUseCase,
	categoryUseCase contract.CategoryUseCase,
	checkoutUseCase contract.CheckoutUseCase,
) *App {
	return &App{
		logger:          logger,
		clientUseCase:   clientUseCase,
		productUseCase:  productUseCase,
		orderUseCase:    orderUseCase,
		paymentUseCase:  paymentUseCase,
		categoryUseCase: categoryUseCase,
		checkoutUseCase: checkoutUseCase,
	}
}

func (e *App) Run(ctx context.Context) error {
	router := mux.NewRouter()

	clientController := controller.NewClientController(e.clientUseCase, e.logger)
	router.HandleFunc("/client", clientController.CreateClient).Methods("POST")
	router.HandleFunc("/client", clientController.GetClientByCpf).Methods("GET")

	productController := controller.NewProductController(e.productUseCase, e.categoryUseCase, e.logger)
	router.HandleFunc("/product", productController.CreateProduct).Methods("POST")
	router.HandleFunc("/category/{categoryId:[0-9]+}/product", productController.GetProductByCategory).Methods("GET")
	router.HandleFunc("/product/{productId:[0-9]+}", productController.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{productId:[0-9]+}", productController.DeleteProduct).Methods("DELETE")

	orderController := controller.NewOrderController(e.orderUseCase, e.productUseCase, e.logger)
	router.HandleFunc("/order", orderController.FindOrders).Methods("GET")
	router.HandleFunc("/order/{orderId:[0-9]+}", orderController.GetOrderById).Methods("GET")
	router.HandleFunc("/order", orderController.CreateOrder).Methods("POST")
	router.HandleFunc("/order/{orderId:[0-9]+}", orderController.UpdateOrderStatusById).Methods("PATCH")

	paymentController := controller.NewPaymentController(e.paymentUseCase, e.logger)
	router.HandleFunc("/payments", paymentController.CreatePayment).Methods("POST")
	router.HandleFunc("/order/{orderId:[0-9]+}/status-payment", paymentController.GetLastPaymentStatus).Methods("GET")

	return http.ListenAndServe(":8080", router)
}
