package repository

import (
	"context"

	"github.com/francocp12/CleanManager/internal/service/contracts"
	"github.com/francocp12/CleanManager/internal/service/dtos"
	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

func (r *MySQLRepository) GetByID(id int) (*dtos.Product, error) {
	var product dtos.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *MySQLRepository) GetAll() ([]dtos.Product, error) {
	var products []dtos.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *MySQLRepository) Create(ctx context.Context, producto *dtos.Product) error {
	return r.db.WithContext(ctx).Create(producto).Error
}

func (r *MySQLRepository) Update(ctx context.Context, producto *dtos.Product) error {
	return r.db.WithContext(ctx).Save(producto).Error
}

func (r *MySQLRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&dtos.Product{}, id).Error
}

func NewMySQLRepository(db *gorm.DB) contracts.ProductoRepository {
	return &MySQLRepository{db: db}
}
