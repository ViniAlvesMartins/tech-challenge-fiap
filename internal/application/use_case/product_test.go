package use_case

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract/mock"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductUseCase_Create(t *testing.T) {
	t.Run("create product successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		p := entity.Product{
			ProductName: "Test product",
			Description: "Product created to test",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		repo := mock.NewMockProductRepository(ctrl)
		repo.EXPECT().Create(p).Return(p, nil).Times(1)

		productUseCase := NewProductUseCase(repo)
		product, err := productUseCase.Create(p)

		assert.Nil(t, err)
		assert.Equal(t, p, *product)
	})

	t.Run("repository error creating product", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		p := entity.Product{
			ProductName: "Test product",
			Description: "Product created to test",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		repo := mock.NewMockProductRepository(ctrl)
		repo.EXPECT().Create(p).Return(p, expectedErr).Times(1)

		productUseCase := NewProductUseCase(repo)
		_, err := productUseCase.Create(p)

		assert.ErrorIs(t, err, expectedErr)
	})
}

func TestProductUseCase_Update(t *testing.T) {
	t.Run("update product successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		p := entity.Product{
			ID:          1,
			ProductName: "Updated test product",
			Description: "Product updated to test",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		updatedProduct := p

		repo := mock.NewMockProductRepository(ctrl)
		repo.EXPECT().Update(p).Return(&updatedProduct, nil).Times(1)

		productUseCase := NewProductUseCase(repo)
		product, err := productUseCase.Update(p, p.ID)

		assert.Nil(t, err)
		assert.Equal(t, updatedProduct, *product)
	})

	t.Run("repository error updating  product", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		p := entity.Product{
			ID:          1,
			ProductName: "Updated test product",
			Description: "Product updated to test",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		repo := mock.NewMockProductRepository(ctrl)
		repo.EXPECT().Update(p).Return(nil, expectedErr).Times(1)

		productUseCase := NewProductUseCase(repo)
		_, err := productUseCase.Update(p, p.ID)

		assert.ErrorIs(t, err, expectedErr)
	})
}

func TestProductUseCase_Delete(t *testing.T) {
	t.Run("delete product successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		p := entity.Product{
			ID:          1,
			ProductName: "Updated test product",
			Description: "Product updated to test",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		repo := mock.NewMockProductRepository(ctrl)
		repo.EXPECT().Delete(p.ID).Return(nil).Times(1)

		productUseCase := NewProductUseCase(repo)
		err := productUseCase.Delete(p.ID)

		assert.Nil(t, err)
	})

	t.Run("repository error deleting product", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		p := entity.Product{
			ID:          1,
			ProductName: "Updated test product",
			Description: "Product updated to test",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		repo := mock.NewMockProductRepository(ctrl)
		repo.EXPECT().Delete(p.ID).Return(expectedErr).Times(1)

		productUseCase := NewProductUseCase(repo)
		err := productUseCase.Delete(p.ID)

		assert.Error(t, expectedErr, err)
	})
}

func TestProductUseCase_GetProductByCategory(t *testing.T) {
	t.Run("list products by category successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		pp := []entity.Product{
			{
				ID:          1,
				ProductName: "Product 1",
				Description: "Product of category 1",
				Price:       123.45,
				CategoryId:  1,
				Active:      true,
			},
			{
				ID:          2,
				ProductName: "Product 2",
				Description: "Product of category 2",
				Price:       123.45,
				CategoryId:  2,
				Active:      true,
			}, {
				ID:          3,
				ProductName: "Product 3",
				Description: "Product of category 1",
				Price:       123.45,
				CategoryId:  1,
				Active:      true,
			},
		}

		expectedProducts := append(pp[:1], pp[2:]...)
		repo := mock.NewMockProductRepository(ctrl)
		repo.EXPECT().GetProductByCategory(1).Return(expectedProducts, nil).Times(1)

		productUseCase := NewProductUseCase(repo)
		products, err := productUseCase.GetProductByCategory(1)

		assert.Nil(t, err)
		assert.Equal(t, expectedProducts, products)
	})

	t.Run("repository error listing products by category", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		repo := mock.NewMockProductRepository(ctrl)
		repo.EXPECT().GetProductByCategory(1).Return(nil, expectedErr).Times(1)

		productUseCase := NewProductUseCase(repo)
		_, err := productUseCase.GetProductByCategory(1)

		assert.ErrorIs(t, err, expectedErr)
	})
}
