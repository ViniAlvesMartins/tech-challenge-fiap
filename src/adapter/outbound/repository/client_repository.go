package repository

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"gorm.io/gorm"
	"log/slog"
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
		return client, errors.New("create client from repository has failed")
	}

	return client, nil
}

func (c *ClientRepository) GetClientByCpf(cpf int) (*entity.Client, error) {

	var client entity.Client

	if result := c.db.Find(&client, "cpf=?", cpf); result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		c.logger.Error("result.Error")
		return nil, errors.New("an error occurred from repository")
	}

	return &client, nil
}
