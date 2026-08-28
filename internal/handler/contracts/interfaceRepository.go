package contracts

import "CleanManager/internal/handler/DTOs/productDTO"

type ProductoRepository interface {
	GetAll() ([]productDTO.Product, error)
	GetByID(id int) (*productDTO.Product, error)
	Create(producto *productDTO.Product) error
	Update(producto *productDTO.Product) error
	Delete(id int) error
}
