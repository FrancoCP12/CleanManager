package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/francocp12/CleanManager/internal/service/dtos"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newMockRepository(t *testing.T) (*MySQLRepository, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening mock database: %v", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("error opening gorm database: %v", err)
	}

	return &MySQLRepository{db: gormDB}, mock
}

func sampleProduct() *dtos.Product {
	return &dtos.Product{
		ID:           1,
		Name:         "Detergente",
		SKU:          "DET-001",
		Price:        12.5,
		CurrentStock: 30,
	}
}
