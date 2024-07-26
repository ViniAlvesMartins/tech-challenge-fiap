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
	result := c.db.Model(&entity.Client{}).Where("cpf = ?", cpf).Update("active", false)
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

	if result := c.db.First(&client, "cpf=? and active = ?", cpf, true); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &client, nil
}

func (c *ClientRepository) GetById(id *int) (*entity.Client, error) {
	var client entity.Client

	if result := c.db.First(&client, "id=? and active = ?", id, true); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &client, nil
}

func (c *ClientRepository) GetByCpfOrEmail(cpf int, email string) (*entity.Client, error) {
	var client entity.Client

	if result := c.db.First(&client, "(cpf=? OR email=?) and active = ?", cpf, email, true); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &client, nil
}
