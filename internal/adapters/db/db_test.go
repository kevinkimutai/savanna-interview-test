package db

import (
	"testing"

	"github.com/kevinkimutai/savanna-app/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {

	db := NewDB(config.DBURL)

	assert.NotNil(t, db, "DBAdapter should not be nil")
}
