package contracts

import (
	"context"

	"github.com/francocp12/CleanManager/internal/service/dtos"
)

type ProductoRepository interface {
	GetAll() ([]dtos.Product, error)
	GetByID(id int) (*dtos.Product, error)
	Create(ctx context.Context, producto *dtos.Product) error
	Update(ctx context.Context, producto *dtos.Product) error
	Delete(ctx context.Context, id int) error
}
