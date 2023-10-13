package repository

import (
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"fmt"
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

func (repo *ClientRepository) Create(client domain.Client) (domain.Client, error) {

	if result := repo.db.Create(&client); result.Error != nil {
		fmt.Println("result.Error")

		fmt.Println(result.Error)
	}

	return client, nil
}
