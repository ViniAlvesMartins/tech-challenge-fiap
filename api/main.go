package main

import (
	"context"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/external/handler/sqs"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/external/service"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/use_case"
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

	sqsService := service.NewSqsService()

	clientRepository := repository.NewClientRepository(db, logger)
	clientUseCase := use_case.NewClientUseCase(clientRepository, logger)

	productRepository := repository.NewProductRepository(db, logger)
	productUseCase := use_case.NewProductUseCase(productRepository, logger)

	orderRepository := repository.NewOrderRepository(db, logger)
	orderUseCase := use_case.NewOrderUseCase(orderRepository, logger)

	categoryRepository := repository.NewCategoryRepository(db, logger)
	categoryUseCase := use_case.NewCategoryUseCase(categoryRepository, logger)

	fmt.Println("aquisd")

	consumerSqs := sqs.NewSqsConsumer(sqsService, orderUseCase, logger)

	app := http_server.NewApp(logger, clientUseCase, productUseCase, orderUseCase, categoryUseCase)

	go app.Run(ctx)
	go consumerSqs.Run()

	select {}
}

func loadConfig() (infra.Config, error) {
	return infra.NewConfig()
}

func loadLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}
