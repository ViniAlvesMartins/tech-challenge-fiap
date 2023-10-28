package dto

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
)

type ClientDto struct {
	ID    int    `json:"id"`
	Cpf   int    `json:"cpf" validate:"required" error:"Campo cpf é obrigatorio"`
	Name  string `json:"name" validate:"required" error:"Campo nome é obrigatorio"`
	Email string `json:"email" validate:"required" error:"Campo email é obrigatorio"`
}

type ClientFields struct {
	Field   string
	Message string
}

func ConvertClientDtoToDomain(dto ClientDto) entity.Client {
	var client = entity.Client{
		ID:    dto.ID,
		Name:  dto.Name,
		Cpf:   dto.Cpf,
		Email: dto.Email,
	}

	return client
}
