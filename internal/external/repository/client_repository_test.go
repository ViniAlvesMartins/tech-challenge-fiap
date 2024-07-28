package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestClientRepository_Create(t *testing.T) {
	t.Run("create client successfully", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test client",
			Email: "testclient@example.com",
		}

		repo := NewClientRepository(db)
		addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")

		expectedSQL := `INSERT INTO "clients" ("cpf","name","email","id") VALUES ($1,$2,$3,$4)`
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(
			client.Cpf, client.Name, client.Email, client.ID,
		).WillReturnRows(addRow)
		mock.ExpectCommit()

		_, err := repo.Create(client)

		assert.Nil(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("error creating client", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		expectedErr := errors.New("scan error")
		defer sqlDB.Close()

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test client",
			Email: "testclient@example.com",
		}

		repo := NewClientRepository(db)

		expectedSQL := `INSERT INTO "clients" ("cpf","name","email","id") VALUES ($1,$2,$3,$4)`
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(
			client.Cpf, client.Name, client.Email, client.ID,
		).WillReturnError(expectedErr)
		mock.ExpectRollback()

		_, err := repo.Create(client)

		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}

func TestClientRepository_GetClientByCpf(t *testing.T) {
	t.Run("get client by cpf successfully", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewClientRepository(db)
		rows := sqlmock.NewRows([]string{"id", "cpf", "name", "email"}).
			AddRow(1, 12345678900, "Test client", "testclient@example.com")

		expectedSQL := `SELECT * FROM "clients" WHERE cpf=$1 ORDER BY "clients"."id" LIMIT $2`
		mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(12345678900, 1).WillReturnRows(rows)
		client, err := repo.GetByCpf(12345678900)

		assert.IsType(t, entity.Client{}, *client)
		assert.Nil(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("record not found", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewClientRepository(db)
		rows := sqlmock.NewRows([]string{"id", "cpf", "name", "email"})

		expectedSQL := `SELECT * FROM "clients" WHERE cpf=$1 ORDER BY "clients"."id" LIMIT $2`
		mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(12345678900, 1).WillReturnRows(rows)
		client, err := repo.GetByCpf(12345678900)

		assert.Nil(t, client)
		assert.Nil(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("database error", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		expectedErr := errors.New("scanErr")
		defer sqlDB.Close()

		repo := NewClientRepository(db)

		expectedSQL := `SELECT * FROM "clients" WHERE cpf=$1 ORDER BY "clients"."id" LIMIT $2`
		mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(12345678900, 1).WillReturnError(expectedErr)
		client, err := repo.GetByCpf(12345678900)

		assert.Nil(t, client)
		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
