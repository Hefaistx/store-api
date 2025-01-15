package repository

import (
	"database/sql"
	conf "enigma-laundry-app/config"
	m "enigma-laundry-app/model"
)

type TransactionRepository interface {
	CreateTransaction(m.Transaction) (m.Transaction, error)
	GetTransactionById(id int) (m.Transaction, error)
	GetTransactions() ([]m.Transaction, error)
	UpdateTransaction(m.Transaction) (m.Transaction, error)
	DeleteTransaction(id int) error
}

type transactionRepository struct {
	db *sql.DB
}

func (db *transactionRepository) CreateTransaction(transaction m.Transaction, details m.TransactionDetail) (m.Transaction, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return m.Transaction{}, err
	}
	err = db.db.QueryRow(conf.CreateTransactionQuery, transaction.ReceivedBy, transaction.CustomerId, transaction.Created_At, transaction.Updated_At).Scan(&transaction.Id)
	if err != nil {
		tx.Rollback()
		return m.Transaction{}, err
	}
	for _, details := range details {
		err = db.db.QueryRow(conf.CreateTransactionDetailQuery, details.ProductId, details.Quantity, transaction.Id).Scan(&details.Id)
		if err != nil {
			details.Id = transaction.Id
			tx.Rollback()
			return m.Transaction{}, err
		}
	}

	return transaction, nil
}

func (db *transactionRepository) GetTransactionById(id int) (m.Transaction, error) {

	var transaction m.Transaction
	err := db.db.QueryRow(conf.GetTransactionByIdQuery, id).Scan(&transaction.Id, &transaction.ReceivedBy, &transaction.Created_At, &transaction.Updated_At)

	if err != nil {
		return m.Transaction{}, err
	}

	return transaction, nil
}

func (db *transactionRepository) GetTransactionsQuery() ([]m.Transaction, error) {

	var transactions []m.Transaction

	rows, err := db.db.Query(conf.GetTransactionsQuery)

	if err != nil {
		return []m.Transaction{}, err
	}

	for rows.Next() {

		var transaction m.Transaction
		err := rows.Scan(&transaction.Id, &transaction.ReceivedBy, &transaction.Created_At, &transaction.Updated_At, &transaction.Finished_At)

		if err != nil {
			return []m.Transaction{}, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (db *transactionRepository) FinishedTransaction(transaction m.Transaction) error {

	_, err := db.db.Exec(conf.FinishedTransactionQuery, transaction.Id)

	if err != nil {
		return err
	}

	return nil
}

func (db *transactionRepository) DeleteTransactionQuery(id int) error {

	_, err := db.db.Exec(conf.DeleteTransactionQuery, id)

	if err != nil {
		return err
	}

	return nil
}
