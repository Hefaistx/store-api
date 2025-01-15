package model

import (
	"time"
)

type Transaction struct {
	Id          int       `json:"transaction_id"`
	ReceivedBy  string    `json:"received_by"`
	CustomerId  int       `json:"customer_id"`
	Created_At  time.Time `json:"created_at"`
	Updated_At  time.Time `json:"updated_at"`
	Finished_At time.Time `json:"finished_at"`
}
