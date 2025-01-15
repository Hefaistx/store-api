package repository

import (
	"database/sql"
	m "enigma-laundry-app/model"
)

type TransactionDetail interface {
	CreateTransactionDetail(transactionDetail m.TransactionDetail) (m.TransactionDetail, error)
	GetTransactionDetailById(id int) (m.TransactionDetail, error)
	GetTransactionDetails(transactionId int) ([]m.TransactionDetail, error)
	UpdateTransactionDetail(transactionDetail m.TransactionDetail) (m.TransactionDetail, error)
	DeleteTransactionDetail(id int) error
}

type transactionDetailRepository struct {
	db *sql.DB
}

// func CreateTransactionDetail(db *sql.DB) TransactionDetail (m.TransactionDetail, error){

// }
