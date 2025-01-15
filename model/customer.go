package model

type Customer struct {
	CustomerID int    `json:"customer_id" db:"customer_id"`
	Name       string `json:"name" db:"name"`
	Phone      int    `json:"phone" db:"phone"`
}
