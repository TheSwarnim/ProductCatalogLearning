package respository

import (
	"errors"
	"go_microservice_learning_1/models"
)

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type ProductRepository interface {
	FindById(id int) (*models.Product, error)
	Save(product *models.Product) error
	Update(product *models.Product) error
	Delete(id int) error
	FindAll() models.Products
}

func NewProductRepository(productList models.Products) ProductRepository {
	return &ProductRepositoryImpl{productList}
}

type ProductRepositoryImpl struct {
	productList models.Products
}

func (p *ProductRepositoryImpl) FindById(id int) (*models.Product, error) {
	for _, product := range p.productList {
		if product.Id == id {
			return product, nil
		}
	}
	return nil, ErrProductNotFound
}

func (p *ProductRepositoryImpl) Save(product *models.Product) error {
	existingProduct, _ := p.FindById(product.Id)
	if existingProduct != nil {
		return ErrProductAlreadyExists
	}

	product.Id = p.getNextProductId()
	p.productList = append(p.productList, product)
	return nil
}

func (p *ProductRepositoryImpl) Update(product *models.Product) error {
	index, _, err := p.findProductAndIndexById(product.Id)
	if err != nil {
		return err
	}

	p.productList[index] = product
	return nil
}

func (p *ProductRepositoryImpl) Delete(id int) error {
	index, _, err := p.findProductAndIndexById(id)
	if err != nil {
		return err
	}

	p.productList = append(p.productList[:index], p.productList[index+1:]...)
	return nil
}

func (p *ProductRepositoryImpl) FindAll() models.Products {
	return p.productList
}

func (p *ProductRepositoryImpl) getNextProductId() int {
	return len(p.productList) + 1
}

func (p *ProductRepositoryImpl) findProductAndIndexById(id int) (int, *models.Product, error) {
	for i, product := range p.productList {
		if product.Id == id {
			return i, product, nil
		}
	}
	return -1, nil, ErrProductNotFound
}
