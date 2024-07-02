package output

import "github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"

type ClientDto struct {
	ID     int    `json:"id"`
	Cpf    int    `json:"cpf"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

func ClientFromEntity(client entity.Client) ClientDto {
	return ClientDto{
		ID:     client.ID,
		Cpf:    client.Cpf,
		Name:   client.Name,
		Email:  client.Email,
		Active: client.Active,
	}
}
