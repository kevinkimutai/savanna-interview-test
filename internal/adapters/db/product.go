package db

import (
	"math"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/utils"
)

func (db *DBAdapter) CreateProduct(product domain.Product) (domain.Product, error) {
	//convert Float64 To pgtype.numeric
	var numeric pgtype.Numeric
	priceStr := strconv.FormatFloat(product.Price, 'f', -1, 64)

	numeric.Scan(priceStr)

	productParams := queries.CreateProductParams{
		Name:     product.Name,
		Price:    numeric,
		ImageUrl: product.ImageURL,
	}

	prod, error := db.queries.CreateProduct(db.ctx, productParams)
	if error != nil {
		return domain.Product{}, error
	}

	return domain.Product{
		ProductID: int(prod.ProductID),
		Name:      prod.Name,
		Price:     product.Price,
		ImageURL:  prod.ImageUrl,
		CreatedAt: prod.CreatedAt.Time,
	}, nil
}

func (db *DBAdapter) GetAllProducts(prodParams queries.ListProductsParams) (domain.ProductsFetch, error) {

	//Get Products
	products, err := db.queries.ListProducts(db.ctx, prodParams)
	if err != nil {
		return domain.ProductsFetch{}, err

	}

	//Get Count
	count, err := db.queries.CountProducts(db.ctx)
	if err != nil {
		return domain.ProductsFetch{}, err

	}

	//Get Page
	page := utils.GetPage(prodParams.Offset, prodParams.Limit)

	//map struct
	var prods []domain.Product

	for _, item := range products {
		// Convert each RequestItem to Product
		product := domain.Product{
			ProductID: int(item.ProductID),
			Name:      item.Name,
			Price:     utils.ConvertNumericToFloat64(item.Price),
			ImageURL:  item.ImageUrl,
			CreatedAt: item.CreatedAt.Time,
		}
		// Append the struct to the struct array
		prods = append(prods, product)
	}

	return domain.ProductsFetch{
		Page:          page,
		NumberOfPages: uint(math.Ceil(float64(count) / float64(prodParams.Limit))),
		Total:         uint(count),
		Data:          prods,
	}, nil

}
func (db *DBAdapter) GetProduct(productID string) (domain.Product, error) {
	prodID, err := utils.ConvertStringToInt64(productID)

	if err != nil {
		return domain.Product{}, err
	}

	product, err := db.queries.GetProduct(db.ctx, prodID)
	if err != nil {
		return domain.Product{}, err
	}

	return domain.Product{
		ProductID: int(product.ProductID),
		Name:      product.Name,
		Price:     float64(product.Price.Exp),
		ImageURL:  product.ImageUrl,
		CreatedAt: product.CreatedAt.Time,
	}, nil

}
func (db *DBAdapter) DeleteProduct(productID string) error {
	prodID, err := utils.ConvertStringToInt64(productID)
	if err != nil {
		return err
	}

	err = db.queries.DeleteProduct(db.ctx, prodID)
	if err != nil {
		return err
	}

	return nil
}
