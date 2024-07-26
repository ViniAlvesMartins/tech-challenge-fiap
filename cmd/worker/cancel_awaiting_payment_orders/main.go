package main

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/postgres"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/use_case"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/config"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/external/repository"
	"log/slog"
	"os"
)

// orders that are more than one hour on AWAITING_PAYMENT status should be canceled
const orderExpirationThreshold = 1

func main() {
	var err error
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

	logger.Info("starting job")

	if err = cancelExpiredOrders(orderUseCase); err != nil {
		logger.Error("error canceling expired orders", slog.Any("error", err.Error()))
	}

	logger.Info("job successfully executed")

}

func cancelExpiredOrders(o contract.OrderUseCase) error {
	return o.CancelExpiredOrders(orderExpirationThreshold)
}

func loadConfig() (config.Config, error) {
	return config.NewConfig()
}

func loadLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}
