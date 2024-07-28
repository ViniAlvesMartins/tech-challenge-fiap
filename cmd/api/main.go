package main

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/postgres"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/use_case"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/config"
	"log/slog"
	"os"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/external/handler/http_server"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/external/repository"
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

	db, err := postgres.NewConnection(cfg.DatabaseHost, cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseDBName, cfg.DatabasePort, cfg.DatabaseSchema)
	if err != nil {
		logger.Error("error connecting tdo database", err)
		panic(err)
	}

	productRepository := repository.NewProductRepository(db)
	productUseCase := use_case.NewProductUseCase(productRepository)

	orderRepository := repository.NewOrderRepository(db)
	orderUseCase := use_case.NewOrderUseCase(orderRepository, productUseCase)

	clientRepository := repository.NewClientRepository(db)
	clientUseCase := use_case.NewClientUseCase(clientRepository, orderUseCase)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryUseCase := use_case.NewCategoryUseCase(categoryRepository)

	app := http_server.NewApp(logger, clientUseCase, productUseCase, orderUseCase, categoryUseCase)

	app.Run(ctx)
}

func loadConfig() (config.Config, error) {
	return config.NewConfig()
}

func loadLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}
