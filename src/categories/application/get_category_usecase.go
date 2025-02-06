package application

import (
    "errors"
    "holamundo/src/categories/domain"
    "holamundo/src/categories/domain/entities"
    "holamundo/src/categories/infrastructure/repositories"
)

type GetCategoryUseCase struct {
    repo domain.CategoryRepository
}

func NewGetCategoryUseCase(repo domain.CategoryRepository) *GetCategoryUseCase {
    return &GetCategoryUseCase{repo: repo}
}

func (uc *GetCategoryUseCase) Execute(id int32) (*entities.Category, error) {
    if repoImpl, ok := uc.repo.(*repositories.MySQLCategoryRepository); ok {
        return repoImpl.GetByID(id)
    }
    return nil, errors.New("")
}
