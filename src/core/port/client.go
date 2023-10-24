package port

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
)

type ClientRepository interface {
	Create(client domain.Client) (domain.Client, error)
	GetClientByCpf(cpf int) (*domain.Client, error)
}

type ClientService interface {
	GetClientByCpf(cpf int) (*domain.Client, error)
	Create(client domain.Client) (*domain.Client, error)
}
