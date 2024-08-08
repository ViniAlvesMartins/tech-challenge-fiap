package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/stretchr/testify/assert"

	"regexp"
	"testing"
)

func TestProductRepository_Create(t *testing.T) {
	t.Run("create product successfully", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		product := entity.Product{
			ID:          1,
			ProductName: "Test product",
			Description: "Test product",
			Price:       123,
			CategoryId:  1,
			Active:      true,
		}

		repo := NewProductRepository(db)
		addRow := sqlmock.
			NewRows([]string{"id"}).
			AddRow("1")

		expectedSQL := `INSERT INTO "products" ("product_name","description","price","category_id","active","id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(
			product.ProductName, product.Description, product.Price, product.CategoryId, product.Active, product.ID,
		).WillReturnRows(addRow)
		mock.ExpectCommit()

		_, err := repo.Create(product)

		assert.Nil(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
