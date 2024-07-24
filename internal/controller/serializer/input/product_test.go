package input

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductDto_ConvertToEntity(t *testing.T) {
	t.Run("convert dto to entity successfully", func(t *testing.T) {
		dto := ProductDto{
			NameProduct: "Test Product",
			Description: "Test created for test purposes",
			Price:       123.45,
			CategoryId:  1,
		}

		assert.IsType(t, entity.Product{}, dto.ConvertToEntity())
	})
}
