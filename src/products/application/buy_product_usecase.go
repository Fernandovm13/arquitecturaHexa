package application

import (
	"fmt"
	"holamundo/src/core"
	"holamundo/src/products/domain"
)

type BuyProductUseCase struct {
	repo     domain.ProductRepository
	rabbitMQ *core.RabbitMQ
}

func NewBuyProductUseCase(repo domain.ProductRepository, rabbitMQ *core.RabbitMQ) *BuyProductUseCase {
	return &BuyProductUseCase{repo: repo, rabbitMQ: rabbitMQ}
}

func (uc *BuyProductUseCase) Execute(productID int32) error {
	product, err := uc.repo.GetByID(productID)
	if err != nil {
		return fmt.Errorf("error al obtener producto: %w", err)
	}

	// Mensaje modificado
	message := fmt.Sprintf("Tu producto '%s' (ID: %d) está en camino", product.Name, product.ID)

	if err := uc.rabbitMQ.PublishMessage(message); err != nil {
		return fmt.Errorf("error al enviar mensaje a RabbitMQ: %w", err)
	}

	return nil
}
