package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/external/database/postgres"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/external/handler/http_server"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/external/repository"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/infra"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/service"
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
	clientService := service.NewClientService(clientRepository, logger)

	productRepository := repository.NewProductRepository(db, logger)
	productService := service.NewProductService(productRepository, logger)

	orderRepository := repository.NewOrderRepository(db, logger)
	orderService := service.NewOrderService(orderRepository, logger)

	paymentRepository := repository.NewPaymentRepository(db, logger)
	paymentService := service.NewPaymentService(paymentRepository)

	externalPaymentRepository := repository.NewExternalPaymentRepository()
	externalPaymentService := service.NewExternalPayment(externalPaymentRepository)

	checkoutService := service.NewCheckoutService(logger, paymentService, orderService, externalPaymentService)

	categoryRepository := repository.NewCategoryRepository(db, logger)
	categoryService := service.NewCategoryService(categoryRepository, logger)

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
