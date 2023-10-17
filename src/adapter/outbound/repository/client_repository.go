package repository

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
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

func (c *ClientRepository) Create(client domain.Client) (domain.Client, error) {

	if result := c.db.Create(&client); result.Error != nil {
		c.logger.Error("result.Error")
	}

	return client, nil
}
