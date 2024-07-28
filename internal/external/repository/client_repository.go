package repository

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"

	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (c *ClientRepository) DeleteClientByCpf(cpf int) error {
	result := c.db.Where("cpf = ?", cpf).Delete(&entity.Client{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *ClientRepository) Create(client entity.Client) (*entity.Client, error) {
	if result := c.db.Create(&client); result.Error != nil {
		return &client, result.Error
	}

	return &client, nil
}

func (c *ClientRepository) GetByCpf(cpf int) (*entity.Client, error) {
	var client entity.Client

	if result := c.db.First(&client, "cpf=?", cpf); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &client, nil
}

func (c *ClientRepository) GetById(id *int) (*entity.Client, error) {
	var client entity.Client

	if result := c.db.First(&client, "id=?", id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &client, nil
}

func (c *ClientRepository) GetByCpfOrEmail(cpf int, email string) (*entity.Client, error) {
	var client entity.Client

	if result := c.db.First(&client, "cpf=? OR email=?", cpf, email); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &client, nil
}
