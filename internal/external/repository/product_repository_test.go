package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"regexp"
	"testing"
)

func TestProductRepository_Create(t *testing.T) {
	t.Run("create product successfully", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		defer sqlDB.Close()

		product := entity.Product{
			ID:          1,
			NameProduct: "Test product",
			Description: "Test product",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		repo := NewClientRepository(db, logger)
		addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")

		expectedSQL := `INSERT INTO "products" ("id","name_product","description","price", "category_id", "active") VALUES ($1,$2,$3,$4,$5,$6)`
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(
			product.ID, product.NameProduct, product.Description, product.Price, product.CategoryId, product.Active,
		).WillReturnRows(addRow)
		mock.ExpectCommit()

		_, err := repo.Create(client)

		assert.Nil(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
