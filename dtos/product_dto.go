package dtos

type ProductDTO struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
