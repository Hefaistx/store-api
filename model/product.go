package model

type Product struct {
	ProductID int     `json:"product_id" db:"product_id"`
	Name      string  `json:"name" db:"name"`
	Unit      string  `json:"unit" db:"unit"`
	Price     float64 `json:"price" db:"price"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt string  `json:"updated_at" db:"updated_at"`
}
