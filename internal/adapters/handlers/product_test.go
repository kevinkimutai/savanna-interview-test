package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupApp() *fiber.App {
	app := fiber.New()
	return app
}

func TestCreateProduct(t *testing.T) {
	app := setupApp()
	mockAPI := new(MockProductApiPort)
	productService := NewProductService(mockAPI)
	app.Post("/products", productService.CreateProduct)

	newProduct := domain.Product{
		Name:     "Test Product",
		Price:    10.50,
		ImageURL: "http://example.com/image.png",
	}
	mockAPI.On("CreateProduct", newProduct).Return(newProduct, nil)

	body, _ := json.Marshal(newProduct)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response domain.ProductResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, newProduct.Name, response.Data.Name)
	assert.Equal(t, newProduct.Price, response.Data.Price)
	assert.Equal(t, newProduct.ImageURL, response.Data.ImageURL)
}

func TestGetAllProducts(t *testing.T) {
	app := setupApp()
	mockAPI := new(MockProductApiPort)
	productService := NewProductService(mockAPI)
	app.Get("/products", productService.GetAllProducts)

	params := domain.ProductParams{}
	productsFetch := domain.ProductsFetch{
		Page:          1,
		NumberOfPages: 1,
		Total:         1,
		Data: []domain.Product{
			{
				Name:     "Test Product",
				Price:    10.50,
				ImageURL: "http://example.com/image.png",
			},
		},
	}
	mockAPI.On("GetAllProducts", params).Return(productsFetch, nil)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response domain.ProductsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, productsFetch.Page, response.Page)
	assert.Equal(t, productsFetch.Total, response.Total)
	assert.Equal(t, productsFetch.Data[0].Name, response.Data[0].Name)
	assert.Equal(t, productsFetch.Data[0].Price, response.Data[0].Price)
	assert.Equal(t, productsFetch.Data[0].ImageURL, response.Data[0].ImageURL)
}

func TestGetProductByID(t *testing.T) {
	app := setupApp()
	mockAPI := new(MockProductApiPort)
	productService := NewProductService(mockAPI)
	app.Get("/products/:productID", productService.GetProductByID)

	productID := "1"
	product := domain.Product{
		Name:     "Test Product",
		Price:    10.50,
		ImageURL: "http://example.com/image.png",
	}
	mockAPI.On("GetProduct", productID).Return(product, nil)

	req := httptest.NewRequest(http.MethodGet, "/products/"+productID, nil)

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response domain.ProductResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, product.Name, response.Data.Name)
	assert.Equal(t, product.Price, response.Data.Price)
	assert.Equal(t, product.ImageURL, response.Data.ImageURL)
}

func TestDeleteProduct(t *testing.T) {
	app := setupApp()
	mockAPI := new(MockProductApiPort)
	productService := NewProductService(mockAPI)
	app.Delete("/products/:productID", productService.DeleteProduct)

	productID := "1"
	mockAPI.On("DeleteProduct", productID).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/products/"+productID, nil)

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
