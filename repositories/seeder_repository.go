package repositories

import (
	"echo-demo-project/models"

	"gorm.io/gorm"
)

type SeederRepositoryQ interface {
	GetSeederByEmail(user *models.Seeder, email string)
}

type SeederRepository struct {
	DB *gorm.DB
}

func NewSeederRepository(db *gorm.DB) *SeederRepository {
	return &SeederRepository{DB: db}
}
