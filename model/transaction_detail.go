package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionDetail struct {
	Id             uuid.UUID `json: id`
	Transaction_Id uuid.UUID `json: transaction_id`
	UserId         uuid.UUID `json: user_id`
	Service_Id     int       `json: "service_id"`
	Quantity       int       `json: quantity`
	Created_At     time.Time `json: created_at`
	Updated_At     time.Time `json: updated_at`
}
