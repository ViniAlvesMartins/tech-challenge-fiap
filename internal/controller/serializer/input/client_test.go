package input

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientDto_ConvertEntity(t *testing.T) {
	t.Run("convert dto to entity successfully", func(t *testing.T) {
		c := ClientDto{
			Cpf:   12345678911,
			Name:  "Test Client",
			Email: "testclient@example.com",
		}

		assert.IsType(t, entity.Client{}, c.ConvertEntity())
	})
}
