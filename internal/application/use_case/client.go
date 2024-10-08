package use_case

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
)

type ClientUseCase struct {
	clientRepository contract.ClientRepository
	orderUseCase     contract.OrderUseCase
}

func NewClientUseCase(c contract.ClientRepository, o contract.OrderUseCase) *ClientUseCase {
	return &ClientUseCase{
		clientRepository: c,
		orderUseCase:     o,
	}
}

func (c *ClientUseCase) Create(client entity.Client) (*entity.Client, error) {
	newClient, err := c.clientRepository.Create(client)
	if err != nil {
		return nil, err
	}

	return newClient, nil
}

func (c *ClientUseCase) GetByCpf(cpf int) (*entity.Client, error) {
	client, err := c.clientRepository.GetByCpf(cpf)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *ClientUseCase) GetById(id *int) (*entity.Client, error) {
	client, err := c.clientRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return client, nil

}

func (c *ClientUseCase) DeleteClientByCpf(cpf int) error {
	client, err := c.clientRepository.GetByCpf(cpf)
	if err != nil {
		return nil
	}

	if err = c.orderUseCase.AnonymizeOrderClient(client.ID); err != nil {
		return nil
	}

	return c.clientRepository.DeleteClientByCpf(cpf)
}

func (c *ClientUseCase) GetByCpfOrEmail(cpf int, email string) (*entity.Client, error) {
	client, err := c.clientRepository.GetByCpfOrEmail(cpf, email)
	if err != nil {
		return nil, err
	}

	return client, nil
}
