package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"regexp"
	"testing"
)

func TestCategoryRepository_GetById(t *testing.T) {
	t.Run("get category by id successfully", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		defer sqlDB.Close()

		repo := NewCategoryRepository(db, logger)
		categories := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Test Category")

		expectedSQL := `SELECT * FROM "categories" WHERE id= $1 ORDER BY "categories"."id" LIMIT 1`
		mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(1).WillReturnRows(categories)
		_, err := repo.GetById(1)

		assert.Nil(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("record not found", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		defer sqlDB.Close()

		repo := NewCategoryRepository(db, logger)
		categories := sqlmock.NewRows([]string{"id", "name"})

		expectedSQL := `SELECT * FROM "categories" WHERE id= $1 ORDER BY "categories"."id" LIMIT 1`
		mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(1).WillReturnRows(categories)
		_, err := repo.GetById(1)

		assert.Nil(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("database error", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		expectedErr := errors.New("scanErr")
		defer sqlDB.Close()

		repo := NewCategoryRepository(db, logger)

		expectedSQL := `SELECT * FROM "categories" WHERE id= $1 ORDER BY "categories"."id" LIMIT 1`
		mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(1).WillReturnError(expectedErr)
		_, err := repo.GetById(1)

		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
