package main

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/postgres"
	sqsservice "github.com/ViniAlvesMartins/tech-challenge-fiap-common/sqs"
	usecase "github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/use_case"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/config"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/external/repository"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/external/service/sqs"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var err error
	var ctx, cancel = context.WithCancel(context.Background())
	var logger = loadLogger()

	logger.Info("Initializing worker...")
	cfg, err := loadConfig()

	if err != nil {
		logger.Error("error loading config", err)
		panic(err)
	}

	db, err := postgres.NewConnection(cfg.DatabaseHost, cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseDBName, cfg.DatabasePort, cfg.DatabaseSchema)
	if err != nil {
		logger.Error("error connecting to database", err)
		panic(err)
	}

	if err != nil {
		logger.Error("error loading config", err)
		panic(err)
	}

	consumer, err := sqsservice.NewConnection(ctx, cfg.OrderStatusQueue, 1, 20)
	if err != nil {
		logger.Error("error connecting to sqs", err)
		panic(err)
	}

	productsRepository := repository.NewProductRepository(db)
	productsUseCase := usecase.NewProductUseCase(productsRepository)

	ordersRepository := repository.NewOrderRepository(db)
	ordersUseCase := usecase.NewOrderUseCase(ordersRepository, productsUseCase)

	orderStatusHandler := sqs.NewOrderStatusUpdateHandler(ordersUseCase, logger)

	logger.Info("Starting consumer...")
	orderStatusUpdateConsumer := sqs.NewConsumer(consumer, orderStatusHandler, logger)

	var wg sync.WaitGroup
	wg.Add(1)
	go orderStatusUpdateConsumer.Start(ctx, &wg)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	cancel()
	wg.Wait()
	logger.Info("Finishing worker...")
}

func loadConfig() (config.Config, error) {
	return config.NewConfig()
}

func loadLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}
