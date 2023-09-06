package services

import "go_microservice_learning_1/models"

type ProductService interface {
	GetProductById(id int) (models.Product, error)
	AddProduct(product models.Product) error
	RemoveProduct(id int) error
	EditProduct(product models.Product) error
	GetAllProducts() (models.Products, error)
}
