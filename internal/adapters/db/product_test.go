package db

import (
	"log"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var PRODUCT domain.Product

func TestCreateProduct(t *testing.T) {
	db := SetupTestDB()

	t.Run("Create Product", func(t *testing.T) {
		product := domain.Product{
			Name:     "test1",
			Price:    100,
			ImageURL: "test",
		}

		prod, err := db.CreateProduct(product)

		PRODUCT = prod

		require.NoError(t, err)
		require.NotNil(t, prod)

		assert.Equal(t, prod.Name, product.Name)
		assert.Equal(t, prod.Price, product.Price)
		assert.Equal(t, prod.ImageURL, product.ImageURL)

	})
}

func TestCreateMultipleProducts(t *testing.T) {
	db := SetupTestDB()
	rand.Seed(time.Now().UnixNano())

	numProducts := 10 // Number of products to create

	for i := 0; i < numProducts; i++ {
		// Generate random product data
		product := generateRandomProduct()

		// Create product
		prod, err := db.CreateProduct(product)
		if err != nil {
			t.Errorf("Failed to create product %d: %v", i+1, err)
			continue // Skip to next iteration on error
		}

		// Validate the returned product
		assert.NotNil(t, prod)
		assert.Equal(t, prod.Name, product.Name)
		assert.Equal(t, prod.Price, product.Price)
		assert.Equal(t, prod.ImageURL, product.ImageURL)
	}
}

// Function to generate random product data
func generateRandomProduct() domain.Product {
	return domain.Product{
		Name:     getRandomString(8),
		Price:    round(rand.Float64()*1000, 2), // Random float64 price between 0 and 1000, rounded to 2 decimals
		ImageURL: "https://example.com/image",   // Example image URL
	}
}

// Function to round a float64 value to the specified number of decimal places
func round(num float64, places int) float64 {
	pow := math.Pow(10, float64(places))
	return math.Round(num*pow) / pow
}

// Function to generate random string of given length
func getRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func TestGetAllProducts(t *testing.T) {
	db := SetupTestDB()

	t.Run("Get Products", func(t *testing.T) {
		params := domain.ProductParams{}
		prodParams := utils.GetProductAPIParams(params)

		productsFetched, err := db.GetAllProducts(prodParams)

		require.NoError(t, err)
		require.NotNil(t, productsFetched)

		//assert.Equal(t, productsFetched.Page, 1)
	})

	t.Run("Get Products With Search Params", func(t *testing.T) {
		params := domain.ProductParams{
			Search: PRODUCT.Name,
		}

		prodParams := utils.GetProductAPIParams(params)

		productsFetched, err := db.GetAllProducts(prodParams)

		require.NoError(t, err)
		require.NotNil(t, productsFetched)

		//assert.Equal(t, productsFetched.Page, 1)
		assert.Equal(t, productsFetched.Data[0].Name, "test1")
	})
}

func TestGetProduct(t *testing.T) {
	db := SetupTestDB()

	t.Run("Get Product With Valid ID", func(t *testing.T) {
		validProductID := PRODUCT.ProductID
		product, err := db.GetProduct(strconv.Itoa(validProductID))

		expectedProduct := PRODUCT

		require.NoError(t, err)
		require.NotNil(t, product)
		assert.Equal(t, product.ProductID, PRODUCT.ProductID)
		assert.Equal(t, product.Name, expectedProduct.Name)
		assert.Equal(t, product.Price, expectedProduct.Price)
	})

	t.Run("Get Product With Invalid ID", func(t *testing.T) {
		invalidProductID := "0"
		_, err := db.GetProduct(invalidProductID)

		require.Error(t, err)
		log.Println(err.Error())

		if err != nil {
			require.Contains(t, err.Error(), "no rows in result set")
		}
	})
}

func TestDeleteProduct(t *testing.T) {
	db := SetupTestDB()
	defer TeardownTestDB(db)

	t.Run("Delete Product With ID", func(t *testing.T) {
		validProductID := PRODUCT.ProductID

		err := db.DeleteProduct(strconv.Itoa(validProductID))
		require.NoError(t, err)
	})

	//Try to fetch deleted product
	t.Run("Get Product With Valid ID", func(t *testing.T) {
		validProductID := PRODUCT.ProductID
		_, err := db.GetProduct(strconv.Itoa(validProductID))

		require.Error(t, err)
		log.Println(err.Error())

		if err != nil {
			require.Contains(t, err.Error(), "no rows in result set")
		}
	})
}
