package repository

import (
	"ecommerce/models"

	"gorm.io/gorm"
)

// ProductRepository interface defines the methods to be implemented
type ProductRepository interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id uint) (*models.Product, error)
}

// GormProductRepository implements ProductRepository with GORM
type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{db: db}
}

func (repo *GormProductRepository) Create(product *models.Product) error {
	return repo.db.Create(product).Error
}

func (repo *GormProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := repo.db.Find(&products).Error
	return products, err
}

func (repo *GormProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := repo.db.First(&product, id).Error
	return &product, err
}
