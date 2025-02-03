package model

import (
	"time"
)

type TransactionDetail struct {
	Id             int       `json: id`
	Transaction_Id int       `json: transaction_id`
	UserId         int       `json: user_id`
	Product_Id     int       `json: "product_id"`
	Quantity       int       `json: quantity`
	Created_At     time.Time `json: created_at`
	Updated_At     time.Time `json: updated_at`
}
