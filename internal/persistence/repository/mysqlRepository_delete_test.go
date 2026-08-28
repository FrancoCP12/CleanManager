package repository

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDelete_Success(t *testing.T) {
	repo, mock := newMockRepository(t)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `products` WHERE `products`.`id` = ?")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.Delete(context.Background(), 1)
	require.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete_ReturnsErrorOnDBFailure(t *testing.T) {
	repo, mock := newMockRepository(t)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `products` WHERE `products`.`id` = ?")).
		WithArgs(1).
		WillReturnError(assert.AnError)
	mock.ExpectRollback()

	err := repo.Delete(context.Background(), 1)
	require.Error(t, err)
}
