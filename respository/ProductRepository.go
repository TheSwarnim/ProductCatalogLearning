package respository

import "go_microservice_learning_1/models"

type ProductRepository interface {
	FindById(id int) (models.Product, error)
	Save(product models.Product) error
	Update(product models.Product) error
	Delete(id int) error
	FindAll() (models.Products, error)
}

type ProductInventoryRepository struct {
	productList models.Products
}
