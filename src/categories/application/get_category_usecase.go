package application

import (
    "errors"
    "holamundo/src/categories/domain"
    "holamundo/src/categories/domain/entities"
)

type GetCategoryUseCase struct {
    repo domain.CategoryRepository
}

func NewGetCategoryUseCase(repo domain.CategoryRepository) *GetCategoryUseCase {
    return &GetCategoryUseCase{repo: repo}
}

func (uc *GetCategoryUseCase) Execute(id int32) (*entities.Category, error) {
    category, err := uc.repo.GetByID(id)
    if err != nil {
        return nil, errors.New("categoría no encontrada")
    }
    return category, nil
}
