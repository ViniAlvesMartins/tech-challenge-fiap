package main

import (
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/infra/database/postgres"
	cli "fiappos/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/handler/httpserver"

	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	postgres.MigrationExecute()
	cli.Execute()
}


//re 