package postgres

import (
	"errors"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/infra"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"log/slog"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migration struct {
	cfg    infra.Config
	db     *gorm.DB
	logger *slog.Logger
}

func NewMigration(db *gorm.DB, cfg infra.Config, log *slog.Logger) *Migration {
	return &Migration{
		db:     db,
		cfg:    cfg,
		logger: log,
	}
}

func (m *Migration) CreateSchema() {
	m.logger.Info("Creating schema")
	m.db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", m.cfg.DatabaseSchema))
}

func (m *Migration) Migrate() {
	var err error

	m.logger.Info("Getting migration instance")
	migration, err := m.getMigrationInstance(m.cfg.MigrationsDir)
	if err != nil {
		m.logger.Error("error creating migration instance", slog.Any("error", err))
	}

	m.logger.Info("Running migrations")
	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		m.logger.Error("error executing migration", slog.Any("error", err))
	}

	m.logger.Info("Migrations finished")
}

func (m *Migration) getMigrationInstance(dir string) (*migrate.Migrate, error) {
	driver, err := m.getDriver()

	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", dir), m.cfg.DatabaseDBName, driver)
}

func (m *Migration) getDriver() (database.Driver, error) {
	db, err := m.db.DB()

	if err != nil {
		return nil, err
	}

	return postgres.WithInstance(db, &postgres.Config{
		SchemaName: m.cfg.DatabaseSchema,
	})
}
