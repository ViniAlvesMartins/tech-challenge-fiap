package postgres

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4/database"
	"gorm.io/gorm"
	"log/slog"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/infra"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migration struct {
	cfg infra.Config
	db  *gorm.DB
	log *slog.Logger
}

func NewMigration(db *gorm.DB, cfg infra.Config, log *slog.Logger) *Migration {
	return &Migration{
		db:  db,
		cfg: cfg,
		log: log,
	}
}

func (m *Migration) Execute() {
	var err error

	migration, err := m.getMigrationInstance(m.cfg.MigrationsDir)
	if err != nil {
		m.log.Error("error creating migration instance", err)
	}

	err = migration.Up()
	if err != nil {
		m.log.Error("error executing migration", err)
	}

	srcErr, dbErr := migration.Close()
	if srcErr != nil || dbErr != nil {
		m.log.Error("error closing migration instance", srcErr, dbErr)
	}
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

	return postgres.WithInstance(db, &postgres.Config{})
}
