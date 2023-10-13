package client

import (
	"errors"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type GetClient struct {
	clientRepository port.ClientRepository
}

func NewClientGetService(clientRepository port.ClientRepository) *GetClient {
	return &GetClient{
		clientRepository: clientRepository,
	}
}

func (srv *GetClient) Get(id string) (domain.Client, error) {
	client, err := srv.clientRepository.Get(id)
	if err != nil {
		return domain.Client{}, errors.New("get client from repository has failed")
	}

	return client, nil
}
