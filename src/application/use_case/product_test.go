package use_case

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract/mock"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

func TestProductUseCase_Create(t *testing.T) {
	t.Run("create product successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

		p := entity.Product{
			NameProduct: "Test product",
			Description: "Product created to test",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		repo := mock.NewMockProductRepository(ctrl)
		repo.EXPECT().Create(p).Return(p, nil).Times(1)

		productUseCase := NewProductUseCase(repo, logger)
		product, err := productUseCase.Create(p)

		assert.Nil(t, err)
		assert.Equal(t, p, *product)
	})

	t.Run("repository error creating procuct", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		expectedErr := errors.New("error connecting to database")

		p := entity.Product{
			NameProduct: "Test product",
			Description: "Product created to test",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		repo := mock.NewMockProductRepository(ctrl)
		repo.EXPECT().Create(p).Return(p, expectedErr).Times(1)

		productUseCase := NewProductUseCase(repo, logger)
		_, err := productUseCase.Create(p)

		assert.ErrorIs(t, err, expectedErr)
	})
}
