package config

//USER QUERY
var CreateUserQuery = "INSERT INTO users(username, password, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING id"
var UpdateUserQuery = "UPDATE Users SET name = $2, phone = $3, address = $4, updated_at = %5 WHERE id = $1"
var DeleteUserQuery = "DELETE FROM users WHERE id = $1"
var GetUserByIdQuery = "SELECT * FROM users WHERE id = $1"
var GetUsersQuery = "SELECT * FROM users"

//PRODUCT QUERY
var AddProductQuery = "INSERT INTO products (name, unit, price, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING id"
var UpdateProductQuery = "UPDATE products SET name = $2, unit = $3, price = $4, updated_at = $5 WHERE id = $1"
var DeleteProductQuery = "DELETE FROM products WHERE id = $1"
var GetProductByIdQuery = "SELECT * FROM products WHERE id = $1"
var GetProductsQuery = "SELECT * FROM products"

//TRANSACTION QUERY
var CreateTransactionQuery = "INSERT INTO Transactions(received_by, created_at, updated_at) VALUES($1, $2, %3) RETURNING id"
var FinishedTransactionQuery = "UPDATE Transactions SET finished_at = $2 WHERE id = $1"
var DeleteTransactionQuery = "DELETE FROM Transactions WHERE id = $1 ON CASCADE"
var GetTransactionByIdQuery = "SELECT transaction_id, received_by, created_at, updated_at, finished_at FROM Transactions WHERE id = $1"
var GetTransactionsQuery = "SELECT transaction_id, received_by, created_at, updated_at, finished_at FROM Transactions"

//TRANSACTION DETAIL QUERY
var CreateTransactionDetailQuery = "INSERT INTO Transaction_Detail(transaction_id, user_id, product_id, quantity, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id"
var UpdateTransactionDetail = "UPDATE Transaction_Detail SET quantity = $2, updated_at = $3 WHERE id = $1"
var DeleteTransactionDetail = "DELETE FROM Transaction_Detail WHERE id = $1"
var GetTransactionDetailById = "SELECT * FROM Transaction_Detail WHERE id = $1"
var GetTransactionDetails = "SELECT * FROM Transaction_Detail WHERE transaction_id = $1"
