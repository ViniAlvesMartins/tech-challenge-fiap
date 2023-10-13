package port

import "fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"

type ClientRepository interface {
	Get(id string) (domain.Client, error)
	Create(client domain.Client) error
}

type ClientService interface {
	Get(id string) (domain.Client, error)
	Create(client domain.Client) error
}
