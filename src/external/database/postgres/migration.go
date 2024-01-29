package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/infra"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrationExecute(cfg *infra.Config, log *slog.Logger) {
	var err error
    var db *sql.DB

    var validationDB = true

    connStr := fmt.Sprintf("host=%s user=%s sslmode=disable password=%s dbname=%s",
        cfg.DatabaseHost, cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseDBName)

    for validationDB {
        db, err = sql.Open("postgres", connStr)

        if err != nil {
            log.Error("error opening postgres connection", err)
            time.Sleep(10 * time.Second)
        } else {
            validationDB = false
        }

    }

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Error("error instantiating postgres driver", err)
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", cfg.MigrationsDir), cfg.DatabaseDBName, driver)
	if err != nil {
		log.Error("error creating migration instance", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Error("error executing migration", err)
	}

	srcErr, dbErr := m.Close()
	if srcErr != nil || dbErr != nil {
		log.Error("error closing migration instance", srcErr, dbErr)
	}
}
