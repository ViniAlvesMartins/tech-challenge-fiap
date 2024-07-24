package main

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/postgres"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/config"
	"log/slog"
	"os"
)

func main() {
	var err error
	var logger = slog.New(slog.NewTextHandler(os.Stderr, nil))

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error("error loading config", err)
		panic(err)
	}

	db, err := postgres.NewConnection(cfg.DatabaseHost, cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseDBName, cfg.DatabasePort, cfg.DatabaseSchema)
	if err != nil {
		logger.Error("error connecting tdo database", err)
		panic(err)
	}

	migration := postgres.NewMigration(db, cfg.DatabaseDBName, cfg.DatabaseSchema, cfg.MigrationsDir, logger)
	migration.CreateSchema()
	migration.Migrate()
}
