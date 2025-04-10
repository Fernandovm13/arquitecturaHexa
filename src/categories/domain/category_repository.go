package domain

import "holamundo/src/categories/domain/entities"

type CategoryRepository interface {
	Save(category *entities.Category) error
	GetAll() ([]entities.Category, error)
	Update(category *entities.Category) error
	Delete(id int32) error
	GetByID(id int32) (*entities.Category, error)
}
