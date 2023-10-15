package main

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/infra"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/infra/database/postgres"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/handler/httpserver"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func main() {
	var err error
	var ctx = context.Background()
	var logger = loadLogger()

	cfg, err := loadConfig()

	if err != nil {
		logger.Error("error loading config", err)
		panic("error loading config")
	}

	db, err := loadDatabase(ctx, cfg)

	httpserver.Run(ctx, logger, db)
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
