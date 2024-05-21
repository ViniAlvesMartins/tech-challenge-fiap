package repository

import (
	"errors"
	"log/slog"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"

	"gorm.io/gorm"
)

type ClientRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewClientRepository(db *gorm.DB, logger *slog.Logger) *ClientRepository {
	return &ClientRepository{
		db:     db,
		logger: logger,
	}
}

func (c *ClientRepository) Create(client entity.Client) (entity.Client, error) {
	if result := c.db.Create(&client); result.Error != nil {
		c.logger.Error("result.Error")
		return client, result.Error
	}

	return client, nil
}

func (c *ClientRepository) GetClientByCpf(cpf int) (*entity.Client, error) {
	var client entity.Client

	if result := c.db.First(&client, "cpf=?", cpf); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		c.logger.Error("result.Error")
		return nil, result.Error
	}

	return &client, nil
}

func (c *ClientRepository) GetClientById(id *int) (*entity.Client, error) {
	var client entity.Client

	if result := c.db.First(&client, "id=?", id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		c.logger.Error("result.Error")
		return nil, result.Error
	}

	return &client, nil
}

func (c *ClientRepository) GetAlreadyExists(cpf int, email string) (*entity.Client, error) {

	var client entity.Client

	if result := c.db.First(&client, "cpf=? OR email=?", cpf, email); result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		c.logger.Error("result.Error")
		return nil, result.Error
	}

	return &client, nil
}
