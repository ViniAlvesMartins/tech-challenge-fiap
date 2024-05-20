package output

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientFromEntity(t *testing.T) {
	t.Run("convert entity to dto successfully", func(t *testing.T) {
		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test client",
			Email: "testclient@example.com",
		}

		assert.IsType(t, ClientDto{}, ClientFromEntity(client))
	})
}
