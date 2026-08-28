package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate_Success(t *testing.T) {
	repo, mock := newMockRepository(t)

	product := sampleProduct()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `products` (`name`,`sku`,`price`,`current_stock`,`id`) VALUES (?,?,?,?,?)")).
		WithArgs(product.Name, product.SKU, product.Price, product.CurrentStock, product.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Create(context.Background(), product)
	require.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate_ReturnsErrorOnDBFailure(t *testing.T) {
	repo, mock := newMockRepository(t)

	product := sampleProduct()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `products` (`name`,`sku`,`price`,`current_stock`,`id`) VALUES (?,?,?,?,?)")).
		WithArgs(product.Name, product.SKU, product.Price, product.CurrentStock, product.ID).
		WillReturnError(errors.New("constraint violation"))
	mock.ExpectRollback()

	err := repo.Create(context.Background(), product)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "constraint violation")
}
