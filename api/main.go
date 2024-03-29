package main

import (
	"context"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/use_case"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/external/service"

	"log/slog"
	"os"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/infra"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/external/database/postgres"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/external/handler/http_server"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/external/repository"
)

// @title           Ze Burguer APIs
// @version         1.0
func main() {
	var err error
	var ctx = context.Background()
	var logger = loadLogger()

	cfg, err := loadConfig()

	if err != nil {
		logger.Error("error loading config", err)
		panic(err)
	}

	db, err := postgres.NewConnection(cfg)
	if err != nil {
		logger.Error("error connecting tdo database", err)
		panic(err)
	}

	clientRepository := repository.NewClientRepository(db, logger)
	clientUseCase := use_case.NewClientUseCase(clientRepository, logger)

	productRepository := repository.NewProductRepository(db, logger)
	productUseCase := use_case.NewProductUseCase(productRepository, logger)

	orderRepository := repository.NewOrderRepository(db, logger)
	orderUseCase := use_case.NewOrderUseCase(orderRepository, logger)

	paymentRepository := repository.NewPaymentRepository(db, logger)
	externalPaymentService := service.NewExternalPayment()
	paymentUseCase := use_case.NewPaymentUseCase(paymentRepository, externalPaymentService, logger, orderUseCase)

	categoryRepository := repository.NewCategoryRepository(db, logger)
	categoryUseCase := use_case.NewCategoryUseCase(categoryRepository, logger)

	app := http_server.NewApp(logger, clientUseCase, productUseCase, orderUseCase, paymentUseCase, categoryUseCase)

	err = app.Run(ctx)

	if err != nil {
		logger.Error("error running application", err)
		panic(err)
	}
}

func loadConfig() (infra.Config, error) {
	return infra.NewConfig()
}

func loadLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}
