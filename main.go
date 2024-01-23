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
	clientUseCase := use_case.NewClientUseCase(clientRepository, logger)

	productRepository := repository.NewProductRepository(db, logger)
	productUseCase := use_case.NewProductUseCase(productRepository, logger)

	orderRepository := repository.NewOrderRepository(db, logger)
	orderUseCase := use_case.NewOrderUseCase(orderRepository, logger)

	paymentRepository := repository.NewPaymentRepository(db, logger)
	paymentUseCase := use_case.NewPaymentUseCase(paymentRepository)

	externalPaymentRepository := repository.NewExternalPaymentRepository()
	externalPaymentUseCase := service.NewExternalPayment(externalPaymentRepository)

	checkoutUseCase := use_case.NewCheckoutUseCase(logger, paymentUseCase, orderUseCase, externalPaymentUseCase)

	categoryRepository := repository.NewCategoryRepository(db, logger)
	categoryUseCase := use_case.NewCategoryUseCase(categoryRepository, logger)

	app := http_server.NewApp(logger, clientUseCase, productUseCase, orderUseCase, paymentUseCase, categoryUseCase, checkoutUseCase)

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
