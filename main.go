package main

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/infra"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/infra/database/postgres"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/handler/http_server"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/outbound/repository"
	srvClient "github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/service/client"
	srvOrder "github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/service/order"
	srcProduct "github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/service/product"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func main() {
	var err error
	var ctx = context.Background()
	var logger = loadLogger()

	postgres.MigrationExecute()

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

	clientRepository := repository.NewClientRepository(db, logger)
	clientService := srvClient.NewClientService(clientRepository, logger)

	productRepository := repository.NewProductRepository(db, logger)
	productService := srcProduct.NewProductService(productRepository, logger)

	orderRepository := repository.NewOrderRepository(db, logger)
	orderService := srvOrder.NewOrderService(orderRepository, logger)

	entry := http_server.NewEntry(logger, clientService, productService, orderService)

	err = entry.Run(ctx)

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
