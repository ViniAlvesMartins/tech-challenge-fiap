package client

import (
	"errors"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type CreateClient struct {
	clientRepository port.ClientRepository
}

func NewClienteCreateService(clientRepository port.ClientRepository) *CreateClient {
	return &CreateClient{
		clientRepository: clientRepository,
	}
}

func (srv *CreateClient) Create(client domain.Client) (domain.Client, error) {
	clientNew := domain.Client{
		Id:    client.Id,
		Cpf:   client.Cpf,
		Name:  client.Name,
		Email: client.Email,
	}

	if err := srv.clientRepository.Create(clientNew); err != nil {
		return domain.Client{}, errors.New("create client into repository has failed")
	}

	return client, nil
}
