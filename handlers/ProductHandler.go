package handlers

import (
	"errors"
	"go_microservice_learning_1/models"
	"go_microservice_learning_1/respository"
	"go_microservice_learning_1/services"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var (
	ErrNoPlaceholderFound     = errors.New("no placeholder found")
	ErrInternalServerError    = errors.New("internal server error")
	ErrUnableToGetAllProducts = errors.New("unable to get all products")
	ErrUnableToAddProduct     = errors.New("unable to add product")
	ErrUnableToUpdateProduct  = errors.New("unable to update product")
	ErrUnableToGetProduct     = errors.New("unable to get product")
	ErrUnableToDeleteProduct  = errors.New("unable to delete product")
)

type ProductHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewProductHandler(productService services.ProductService, l *log.Logger) ProductHandler {
	return &ProductHandlerImpl{
		productService: productService,
		l:              l,
	}
}

type ProductHandlerImpl struct {
	productService services.ProductService
	l              *log.Logger
}

func (p *ProductHandlerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id, err := p.getPlaceholder(r, p.l)

		p.l.Println(id, err)
		if err != nil {
			if errors.Is(err, ErrNoPlaceholderFound) {
				p.getAllProducts(w)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		p.getProductById(id, w)
	case http.MethodPost:
		p.addProduct(w, r)
	case http.MethodPut:
		p.editProduct(w, r)
	case http.MethodDelete:
		p.removeProduct(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *ProductHandlerImpl) removeProduct(w http.ResponseWriter, r *http.Request) {
	id, err := p.getPlaceholder(r, p.l)
	if err != nil {
		http.Error(w, ErrUnableToDeleteProduct.Error(), http.StatusInternalServerError)
		return
	}

	err = p.productService.RemoveProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (p *ProductHandlerImpl) editProduct(w http.ResponseWriter, r *http.Request) {
	id, err := p.getPlaceholder(r, p.l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product := &models.Product{}
	err = product.FromJSON(r.Body)

	p.l.Println("product received for update: ", product)
	if err != nil {
		http.Error(w, ErrUnableToUpdateProduct.Error(), http.StatusInternalServerError)
		return
	}

	err = p.productService.EditProduct(id, product)

	if errors.Is(err, respository.ErrProductNotFound) {
		http.Error(w, respository.ErrProductNotFound.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, ErrUnableToUpdateProduct.Error(), http.StatusInternalServerError)
	}
}

func (p *ProductHandlerImpl) addProduct(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{}
	err := product.FromJSON(r.Body)

	if err != nil {
		http.Error(w, ErrUnableToAddProduct.Error(), http.StatusInternalServerError)
		return
	}

	err = p.productService.AddProduct(product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (p *ProductHandlerImpl) getAllProducts(w http.ResponseWriter) {
	productList := p.productService.GetAllProducts()
	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, ErrUnableToGetAllProducts.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *ProductHandlerImpl) getProductById(id int, w http.ResponseWriter) {
	product, err := p.productService.GetProductById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = product.ToJSON(w)
	if err != nil {
		http.Error(w, ErrUnableToGetProduct.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *ProductHandlerImpl) getPlaceholder(r *http.Request, l *log.Logger) (int, error) {
	reg := regexp.MustCompile(`/([0-9]+)`)
	g := reg.FindAllStringSubmatch(r.URL.Path, -1)
	l.Println(g)

	if len(g) != 1 || len(g[0]) != 2 {
		return -1, ErrNoPlaceholderFound
	}

	idString := g[0][1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		return -1, ErrInternalServerError
	}

	return id, nil
}
