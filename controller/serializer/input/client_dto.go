package input

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
)

type ClientDto struct {
	ID    int    `json:"id"`
	Cpf   int    `json:"cpf" validate:"required" error:"Campo cpf é obrigatorio"`
	Name  string `json:"name" validate:"required" error:"Campo nome é obrigatorio"`
	Email string `json:"email" validate:"required" error:"Campo email é obrigatorio"`
}

func (c *ClientDto) ConvertEntity() entity.Client {
	return entity.Client{
		ID:    c.ID,
		Name:  c.Name,
		Cpf:   c.Cpf,
		Email: c.Email,
	}
}
