package db

import (
	"context"
	"testing"

	"github.com/kevinkimutai/savanna-app/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {

	db := NewDB(config.DBURL)

	assert.NotNil(t, db, "DBAdapter should not be nil")
}

var testDBURL = config.DBURL

func SetupTestDB() *DBAdapter {
	db := NewDB(testDBURL)
	return db
}

func TeardownTestDB(db *DBAdapter) {
	db.conn.Exec(context.Background(), "DELETE FROM order_items")
	db.conn.Exec(context.Background(), "DELETE FROM orders")
	db.conn.Exec(context.Background(), "DELETE FROM products")
	db.conn.Exec(context.Background(), "DELETE FROM customers")
	db.conn.Close()
}
