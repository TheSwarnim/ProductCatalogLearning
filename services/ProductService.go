package services

import (
	"go_microservice_learning_1/models"
	"go_microservice_learning_1/respository"
)

type ProductService interface {
	GetProductById(id int) (*models.Product, error)
	AddProduct(product *models.Product) error
	RemoveProduct(id int) error
	EditProduct(product *models.Product) error
	GetAllProducts() (models.Products, error)
}

func NewProductService(productRepository respository.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository}
}

type ProductServiceImpl struct {
	productRepository respository.ProductRepository
}

func (p *ProductServiceImpl) GetProductById(id int) (*models.Product, error) {
	return p.productRepository.FindById(id)
}

func (p *ProductServiceImpl) AddProduct(product *models.Product) error {
	return p.productRepository.Save(product)
}

func (p *ProductServiceImpl) RemoveProduct(id int) error {
	return p.productRepository.Delete(id)
}

func (p *ProductServiceImpl) EditProduct(product *models.Product) error {
	return p.EditProduct(product)
}

func (p *ProductServiceImpl) GetAllProducts() (models.Products, error) {
	return p.GetAllProducts()
}
