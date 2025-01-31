package application

import (
	"holamundo/src/products/domain"       
	"holamundo/src/products/domain/entities" 
)

type CreateProductUseCase struct {
	repo domain.ProductRepository 
}

func NewCreateProductUseCase(repo domain.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{repo: repo}
}

func (uc *CreateProductUseCase) Execute(product *entities.Product) error {
	return uc.repo.Save(product)
}
