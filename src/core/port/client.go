package port

import "fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"

type ClientRepository interface {
	Create(client domain.Client) (domain.Client, error)
}

type ClientService interface {
	Create(id string, cpf int32, name string, email string) (domain.Client, error)
}
