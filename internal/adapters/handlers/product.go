package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/ports"
)

type ProductService struct {
	api ports.ProductApiPort
}

func NewProductService(api ports.ProductApiPort) *ProductService {
	return &ProductService{api: api}
}

func (s *ProductService) CreateProduct(c *fiber.Ctx) error {
	newProduct := domain.Product{}

	//Bind To struct
	if err := c.BodyParser(&newProduct); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Check missing values
	product, err := domain.NewProductDomain(newProduct)
	if err != nil {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 400,
				Message:    err.Error(),
			})
	}

	//Create Product API
	prod, err := s.api.CreateProduct(product)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}

	return c.Status(201).JSON(
		domain.ProductResponse{
			StatusCode: 201,
			Message:    "Product created successfully",
			Data:       prod,
		})
}

func (s *ProductService) GetAllProducts(c *fiber.Ctx) error {
	//Get Query Params
	m := c.Queries()

	//Bind To ProductParams
	prodParams := domain.CheckProductsParams(m)

	//Get All Products API
	prod, err := s.api.GetAllProducts(prodParams)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}
	return c.Status(200).JSON(
		domain.ProductsResponse{
			StatusCode:    200,
			Message:       "Successfully retrieved products",
			Page:          prod.Page,
			NumberOfPages: prod.NumberOfPages,
			Total:         prod.Total,
			Data:          prod.Data,
		})
}

func (s *ProductService) GetProductByID(c *fiber.Ctx) error {
	productID := c.Params("productID")

	//Get Product API
	product, err := s.api.GetProduct(productID)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(
		domain.ProductResponse{
			StatusCode: 200,
			Message:    "Successfully retrieved product",
			Data:       product})
}

func (s *ProductService) DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("productID")

	//Delete Product API
	err := s.api.DeleteProduct(productID)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(204).JSON(
		domain.CustomerResponse{
			StatusCode: 204,
			Message:    "Successfully Deleted product",
		})

}
