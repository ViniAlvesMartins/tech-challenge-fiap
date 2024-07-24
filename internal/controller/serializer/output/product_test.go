package output

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductFromEntity(t *testing.T) {
	t.Run("convert entity to dto successfully", func(t *testing.T) {
		prod := entity.Product{
			ID:          1,
			NameProduct: "Test product",
			Description: "Test product",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		assert.IsType(t, ProductDto{}, ProductFromEntity(prod))
	})
}

func TestProductListFromEntity(t *testing.T) {
	t.Run("convert product list to dto successfully", func(t *testing.T) {
		prods := []entity.Product{
			{
				ID:          1,
				NameProduct: "Test product 1",
				Description: "Test product 1",
				Price:       123.45,
				CategoryId:  1,
				Active:      true,
			}, {
				ID:          2,
				NameProduct: "Test product 2",
				Description: "Test product 2",
				Price:       123.45,
				CategoryId:  1,
				Active:      true,
			},
		}

		dto := ProductListFromEntity(prods)
		assert.Equal(t, 2, len(dto))
	})
}
