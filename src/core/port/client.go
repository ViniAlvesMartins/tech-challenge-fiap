package port

import "github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"

type ClientRepository interface {
	Create(client entity.Client) (entity.Client, error)
	GetClientByCpf(cpf int) (*entity.Client, error)
	GetAlreadyExists(cpf int, email string) (*entity.Client, error)
}

type ClientService interface {
	GetClientByCpf(cpf int) (*entity.Client, error)
	Create(client entity.Client) (*entity.Client, error)
	GetAlreadyExists(cpf int, email string) (*entity.Client, error)
}
