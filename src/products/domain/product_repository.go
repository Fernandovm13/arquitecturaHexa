package domain

import "holamundo/src/products/domain/entities"

type ProductRepository interface {
	Save(product *entities.Product) error
	GetAll() ([]entities.Product, error)
	Update(product *entities.Product) error
	Delete(id int32) error
	GetByID(id int32) (*entities.Product, error)
}
