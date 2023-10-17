package postgres

import (
	"database/sql"
	"fmt"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/infra"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrationExecute() {
	cfg, err := infra.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	connStr := fmt.Sprintf("host=%s user=%s sslmode=disable password=%s",
	cfg.DatabaseHost, cfg.DatabaseUsername, cfg.DatabasePassword)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
        "file://infra/database/migrations",
        "postgres", driver)
	if err != nil {
		fmt.Println(err)
	}
	
	m.Up()
	m.Close()
}