package client

import (
	"errors"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
)

type CreateService struct {
	clientRepository port.ClientRepository
}

func createService(clientRepository port.ClientRepository) *CreateService {
	return &CreateService{
		clientRepository: clientRepository,
	}
}

func (srv *CreateService) Create(id string, cpf int32, name string, email string) (domain.Client, error) {
	clientNew := domain.Client{
		Id:    id,
		Cpf:   cpf,
		Name:  name,
		Email: email,
	}

	client, err := srv.clientRepository.Create(clientNew)

	if err != nil {
		return domain.Client{}, errors.New("create product from repository has failed")
	}

	return client, nil
}
