package application

import (
    "encoding/json"
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

    //  JSON con la información necesaria para el pago
    purchaseData := struct {
        ProductID   int32   `json:"product_id"`
        ProductName string  `json:"product_name"`
        Amount      float32 `json:"amount"`
    }{
        ProductID:   product.ID,
        ProductName: product.Name,
        Amount:      product.Price,
    }

    jsonBytes, err := json.Marshal(purchaseData)
    if err != nil {
        return fmt.Errorf("error al serializar mensaje JSON: %w", err)
    }

    if err := uc.rabbitMQ.PublishMessage(string(jsonBytes)); err != nil {
        return fmt.Errorf("error al enviar mensaje a RabbitMQ: %w", err)
    }

    return nil
}
