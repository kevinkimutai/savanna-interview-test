package db

import (
	"log"
	"testing"

	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateCustomer(t *testing.T) {
	db := NewDB(config.DBURL)
	t.Run("Valid Customer", func(t *testing.T) {
		newCustomer := queries.CreateCustomerParams{
			CustomerID: "1a1a",
			Name:       "test test",
			Email:      "test@test.com",
		}

		cus, err := db.CreateCustomer(newCustomer)

		require.NoError(t, err)
		require.NotNil(t, cus)
		assert.Equal(t, newCustomer.CustomerID, cus.CustomerID)
		assert.Equal(t, newCustomer.Name, cus.Name)
		assert.Equal(t, newCustomer.Email, cus.Email)
	})

	t.Run("Duplicate Email", func(t *testing.T) {
		duplicateCustomer := queries.CreateCustomerParams{
			CustomerID: "1b1b", // Same ID as above
			Name:       "duplicate test",
			Email:      "test@test.com",
		}

		_, err := db.CreateCustomer(duplicateCustomer)

		require.Error(t, err, "expecting error for duplicate email")

		if err != nil {
			require.Contains(t, err.Error(), "duplicate key value violates unique constraint")
		}
	})

}

func TestGetCustomerByEmail(t *testing.T) {
	db := NewDB(config.DBURL)

	t.Run("Valid Email", func(t *testing.T) {
		validCustomer := queries.CreateCustomerParams{
			CustomerID: "1a1a",
			Name:       "test test",
			Email:      "test@test.com",
		}

		cus, err := db.GetCustomerByEmail(validCustomer.Email)

		require.NoError(t, err)
		require.NotNil(t, cus)
		assert.Equal(t, validCustomer.CustomerID, cus.CustomerID)
		assert.Equal(t, validCustomer.Name, cus.Name)
		assert.Equal(t, validCustomer.Email, cus.Email)
	})

	t.Run("Invalid Email", func(t *testing.T) {
		invalidEmail := "invalidemail@test.com"

		_, err := db.GetCustomerByEmail(invalidEmail)

		require.Error(t, err)
		log.Println(err.Error())

		if err != nil {
			require.Contains(t, err.Error(), "no rows in result set")
		}

	})
}
