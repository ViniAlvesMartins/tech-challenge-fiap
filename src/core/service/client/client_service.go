package service

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"log/slog"
)

type CreateService struct {
	clientRepository port.ClientRepository
	logger           *slog.Logger
}

func NewClientService(clientRepository port.ClientRepository, logger *slog.Logger) *CreateService {
	return &CreateService{
		clientRepository: clientRepository,
		logger:           logger,
	}
}

func (c *CreateService) Create(cpf int, name string, email string) (*domain.Client, error) {
	clientNew := domain.Client{
		Cpf:   cpf,
		Name:  name,
		Email: email,
	}

	client, err := c.clientRepository.Create(clientNew)

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *CreateService) GetClientByCpf(cpf int) (*domain.Client, error) {
	client, err := c.clientRepository.GetClientByCpf(cpf)

	if err != nil {
		return nil, err
	}

	return client, nil

}
