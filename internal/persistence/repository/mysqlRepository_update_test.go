package repository

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdate_Success(t *testing.T) {
	repo, mock := newMockRepository(t)

	product := sampleProduct()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `products` SET `name`=?,`sku`=?,`price`=?,`current_stock`=? WHERE `id` = ?")).
		WithArgs(product.Name, product.SKU, product.Price, product.CurrentStock, product.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Update(context.Background(), product)
	require.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate_ReturnsErrorOnDBFailure(t *testing.T) {
	repo, mock := newMockRepository(t)

	product := sampleProduct()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `products` SET `name`=?,`sku`=?,`price`=?,`current_stock`=? WHERE `id` = ?")).
		WithArgs(product.Name, product.SKU, product.Price, product.CurrentStock, product.ID).
		WillReturnError(assert.AnError)
	mock.ExpectRollback()

	err := repo.Update(context.Background(), product)
	require.Error(t, err)
}
