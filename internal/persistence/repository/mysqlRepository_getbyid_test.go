package repository

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetByID_Found(t *testing.T) {
	repo, mock := newMockRepository(t)

	rows := sqlmock.NewRows([]string{"id", "name", "sku", "price", "current_stock"}).
		AddRow(1, "Detergente", "DET-001", 12.5, 30)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `products`.`id` = ? ORDER BY `products`.`id` LIMIT ?")).
		WithArgs(1, 1).
		WillReturnRows(rows)

	product, err := repo.GetByID(1)
	require.NoError(t, err)
	require.NotNil(t, product)

	assert.Equal(t, 1, product.ID)
	assert.Equal(t, "Detergente", product.Name)
	assert.Equal(t, "DET-001", product.SKU)
	assert.Equal(t, 12.5, product.Price)
	assert.Equal(t, 30, product.CurrentStock)
}

func TestGetByID_NotFound(t *testing.T) {
	repo, mock := newMockRepository(t)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `products`.`id` = ? ORDER BY `products`.`id` LIMIT ?")).
		WithArgs(999, 1).
		WillReturnError(sql.ErrNoRows)

	product, err := repo.GetByID(999)
	require.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestGetByID_ReturnsErrorOnDBFailure(t *testing.T) {
	repo, mock := newMockRepository(t)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `products`.`id` = ? ORDER BY `products`.`id` LIMIT ?")).
		WithArgs(1, 1).
		WillReturnError(errors.New("query failed"))

	product, err := repo.GetByID(1)
	require.Error(t, err)
	assert.Nil(t, product)
}
