package entities

type Product struct {
	ID    int32   `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	CategoryID *int32  `json:"category_id"`
}

func NewProduct(name string, price float32) *Product {
	return &Product{Name: name, Price: price}
}
