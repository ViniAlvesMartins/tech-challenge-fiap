package service

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"log/slog"
)

type ClientService struct {
	clientRepository port.ClientRepository
	logger           *slog.Logger
}

func NewClientService(clientRepository port.ClientRepository, logger *slog.Logger) *ClientService {
	return &ClientService{
		clientRepository: clientRepository,
		logger:           logger,
	}
}

func (c *ClientService) Create(client entity.Client) (*entity.Client, error) {

	clientNew, err := c.clientRepository.Create(client)

	if err != nil {
		return nil, err
	}

	return &clientNew, nil
}

func (c *ClientService) GetClientByCpf(cpf int) (*entity.Client, error) {
	client, err := c.clientRepository.GetClientByCpf(cpf)

	if err != nil {
		return nil, err
	}

	return client, nil

}

func (c *ClientService) GetAlreadyExists(cpf int, email string) (*entity.Client, error) {
	client, err := c.clientRepository.GetAlreadyExists(cpf, email)

	if err != nil {
		return nil, err
	}

	return client, nil

}
