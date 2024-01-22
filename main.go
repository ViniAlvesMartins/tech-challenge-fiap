package main

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/application/use_case"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/external/service"

	"log/slog"
	"os"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/external/database/postgres"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/external/handler/http_server"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/external/repository"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/infra"

	"gorm.io/gorm"
)

func main() {
	var err error
	var ctx = context.Background()
	var logger = loadLogger()

	cfg, err := loadConfig()

	if err != nil {
		logger.Error("error loading config", err)
		panic(err)
	}

	db, err := loadDatabase(ctx, cfg)

	if err != nil {
		logger.Error("error connecting tdo database", err)
		panic(err)
	}

	postgres.MigrationExecute(&cfg, logger)

	clientRepository := repository.NewClientRepository(db, logger)
	clientService := use_case.NewClientUseCase(clientRepository, logger)

	productRepository := repository.NewProductRepository(db, logger)
	productService := use_case.NewProductUseCase(productRepository, logger)

	orderRepository := repository.NewOrderRepository(db, logger)
	orderService := use_case.NewOrderUseCase(orderRepository, logger)

	paymentRepository := repository.NewPaymentRepository(db, logger)
	paymentService := use_case.NewPaymentUseCase(paymentRepository)

	externalPaymentRepository := repository.NewExternalPaymentRepository()
	externalPaymentService := service.NewExternalPayment(externalPaymentRepository)

	checkoutService := use_case.NewCheckoutUseCase(logger, paymentService, orderService, externalPaymentService)

	categoryRepository := repository.NewCategoryRepository(db, logger)
	categoryService := use_case.NewCategoryUseCase(categoryRepository, logger)

	app := http_server.NewApp(logger, clientService, productService, orderService, paymentService, categoryService, checkoutService)

	err = app.Run(ctx)

	if err != nil {
		logger.Error("error running application", err)
		panic(err)
	}
}

func loadDatabase(ctx context.Context, cfg infra.Config) (*gorm.DB, error) {
	return postgres.NewConnection(
		ctx,
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
		cfg,
	)
}

func loadConfig() (infra.Config, error) {
	return infra.NewConfig()
}

func loadLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}
