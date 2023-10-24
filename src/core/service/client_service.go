package service

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
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

func (c *ClientService) Create(client domain.Client) (*domain.Client, error) {

	clientNew, err := c.clientRepository.Create(client)

	if err != nil {
		return nil, err
	}

	return &clientNew, nil
}
