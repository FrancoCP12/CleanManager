package repository

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/francocp12/CleanManager/internal/service/dtos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAll_ReturnsProducts(t *testing.T) {
	repo, mock := newMockRepository(t)

	rows := sqlmock.NewRows([]string{"id", "name", "sku", "price", "current_stock"}).
		AddRow(1, "Detergente", "DET-001", 12.5, 30).
		AddRow(2, "Lavandina", "LAV-001", 8.0, 15)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products`")).
		WillReturnRows(rows)

	products, err := repo.GetAll()
	require.NoError(t, err)
	require.Len(t, products, 2)

	assert.Equal(t, []dtos.Product{
		{ID: 1, Name: "Detergente", SKU: "DET-001", Price: 12.5, CurrentStock: 30},
		{ID: 2, Name: "Lavandina", SKU: "LAV-001", Price: 8.0, CurrentStock: 15},
	}, products)
}

func TestGetAll_ReturnsErrorOnDBFailure(t *testing.T) {
	repo, mock := newMockRepository(t)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products`")).
		WillReturnError(errors.New("connection lost"))

	products, err := repo.GetAll()
	require.Error(t, err)
	assert.Nil(t, products)
	assert.Contains(t, err.Error(), "connection lost")
}
