package application

import (
	"holamundo/src/categories/domain"
	"holamundo/src/categories/domain/entities"
)

type CreateCategoryUseCase struct {
	repo domain.CategoryRepository
}

func NewCreateCategoryUseCase(repo domain.CategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{repo: repo}
}

func (uc *CreateCategoryUseCase) Execute(category *entities.Category) error {
	if err := category.EncryptSecret(); err != nil {
		return err
	}
	return uc.repo.Save(category)
}
