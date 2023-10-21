package port

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
)

type ClientRepository interface {
	Create(client domain.Client) (domain.Client, error)
}

type ClientService interface {
	Create(client domain.Client) (*domain.Client, error)
}
