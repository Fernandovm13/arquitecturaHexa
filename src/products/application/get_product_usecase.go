package application

import (
    "errors"
    "holamundo/src/products/domain"
    "holamundo/src/products/domain/entities"
    "holamundo/src/products/infrastructure/repositories"
)

type GetProductUseCase struct {
    repo domain.ProductRepository
}

func NewGetProductUseCase(repo domain.ProductRepository) *GetProductUseCase {
    return &GetProductUseCase{repo: repo}
}

func (uc *GetProductUseCase) Execute(id int32) (*entities.Product, error) {
    if repoImpl, ok := uc.repo.(*repositories.MySQLProductRepository); ok {
        return repoImpl.GetByID(id)
    }
    return nil, errors.New("el repositorio no soporta GetByID")
}
