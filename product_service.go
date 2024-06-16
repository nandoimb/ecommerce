package services

import (
	"ecommerce/models"
	"ecommerce/repository"
)

// ProductService interface defines the methods to be implemented
type ProductService interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
}

// ProductServiceImpl implements ProductService
type ProductServiceImpl struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{repo: repo}
}

func (service *ProductServiceImpl) CreateProduct(product *models.Product) error {
	return service.repo.Create(product)
}

func (service *ProductServiceImpl) GetAllProducts() ([]models.Product, error) {
	return service.repo.GetAll()
}

func (service *ProductServiceImpl) GetProductByID(id uint) (*models.Product, error) {
	return service.repo.GetByID(id)
}
