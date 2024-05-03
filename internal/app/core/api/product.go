package application

import (
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/ports"
	"github.com/kevinkimutai/savanna-app/internal/utils"
)

type ProductRepo struct {
	db ports.ProductRepoPort
}

func NewProductRepo(db ports.ProductRepoPort) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) CreateProduct(product domain.Product) (domain.Product, error) {
	product, err := r.db.CreateProduct(product)

	return product, err
}

func (r *ProductRepo) GetAllProducts(prodParams domain.ProductParams) (domain.ProductsFetch, error) {

	//GetAPIParams
	params := utils.GetProductAPIParams(prodParams)

	data, err := r.db.GetAllProducts(params)

	return data, err
}

func (r *ProductRepo) GetProduct(prodID string) (domain.Product, error) {
	//TODO:HANDLE ERRORS
	product, err := r.db.GetProduct(prodID)

	return product, err
}
func (r *ProductRepo) DeleteProduct(prodID string) error {
	//TODO:HANDLE ERRORS
	err := r.db.DeleteProduct(prodID)

	return err
}
