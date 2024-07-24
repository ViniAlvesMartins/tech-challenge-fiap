package http_server

import (
	"context"
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/doc/swagger"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/ViniAlvesMartins/tech-challenge-fiap/doc/swagger"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/controller/controller"
	"github.com/swaggo/http-swagger/v2"

	"github.com/gorilla/mux"
)

type App struct {
	logger          *slog.Logger
	clientUseCase   contract.ClientUseCase
	productUseCase  contract.ProductUseCase
	orderUseCase    contract.OrderUseCase
	categoryUseCase contract.CategoryUseCase
}

func NewApp(logger *slog.Logger,
	clientUseCase contract.ClientUseCase,
	productUseCase contract.ProductUseCase,
	orderUseCase contract.OrderUseCase,
	categoryUseCase contract.CategoryUseCase,
) *App {
	return &App{
		logger:          logger,
		clientUseCase:   clientUseCase,
		productUseCase:  productUseCase,
		orderUseCase:    orderUseCase,
		categoryUseCase: categoryUseCase,
	}
}

func (e *App) Run(ctx context.Context) {
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
	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)

	swagger.SwaggerInfo.Title = "Ze Burguer Order API"
	swagger.SwaggerInfo.Version = "1.0"

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctxShutdown); err != nil {
		log.Fatal(err)
	}
}
