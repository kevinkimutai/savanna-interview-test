package db

import (
	"context"
	"math"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/utils"
)

func (db *DBAdapter) CreateProduct(product domain.Product) (domain.Product, error) {
	ctx := context.Background()

	//convert Float64 To pgtype.numeric
	var numeric pgtype.Numeric
	priceStr := strconv.FormatFloat(product.Price, 'f', -1, 64)

	numeric.Scan(priceStr)

	productParams := queries.CreateProductParams{
		Name:     product.Name,
		Price:    numeric,
		ImageUrl: product.ImageURL,
	}

	prod, error := db.queries.CreateProduct(ctx, productParams)
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
	ctx := context.Background()

	//Get Products
	products, err := db.queries.ListProducts(ctx, prodParams)
	if err != nil {
		return domain.ProductsFetch{}, err

	}

	//Get Count
	count, err := db.queries.CountProducts(ctx, queries.CountProductsParams{
		Column1: prodParams.Column1,
		Price:   prodParams.Price,
		Price_2: prodParams.Price_2,
	})

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
	ctx := context.Background()

	prodID, err := utils.ConvertStringToInt64(productID)

	if err != nil {
		return domain.Product{}, err
	}

	product, err := db.queries.GetProduct(ctx, prodID)
	if err != nil {
		return domain.Product{}, err
	}

	return domain.Product{
		ProductID: int(product.ProductID),
		Name:      product.Name,
		Price:     utils.ConvertNumericToFloat64(product.Price),
		ImageURL:  product.ImageUrl,
		CreatedAt: product.CreatedAt.Time,
	}, nil

}
func (db *DBAdapter) DeleteProduct(productID string) error {
	ctx := context.Background()

	prodID, err := utils.ConvertStringToInt64(productID)
	if err != nil {
		return err
	}

	err = db.queries.DeleteProduct(ctx, prodID)
	if err != nil {
		return err
	}

	return nil
}
