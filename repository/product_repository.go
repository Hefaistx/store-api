package repository

import (
	"database/sql"
	"fmt"
	conf "tokocikbosapi/config"
	m "tokocikbosapi/model"
)

type productRepository struct {
	db *sql.DB
}

type ProductRepository interface {
	AddProduct(product m.Product) (m.Product, error)
	GetProductById(id int) (m.Product, error)
	GetProducts() ([]m.Product, error)
	UpdateProduct(product m.Product) (m.Product, error)
	DeleteProduct(id int) error
}

func (db *productRepository) AddProduct(product m.Product) (m.Product, error) {
	err := db.db.QueryRow(conf.AddProductQuery, product.Name, product.Unit, product.Price, product.CreatedAt, product.UpdatedAt).Scan(&product.ProductID)
	if err != nil {
		return m.Product{}, fmt.Errorf("error creating service")
	}

	return product, nil
}

func (db *productRepository) GetProductById(id int) (m.Product, error) {
	var product m.Product

	err := db.db.QueryRow(conf.GetProductByIdQuery, id).Scan(&product.ProductID, &product.Name, &product.Unit, &product.Price, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return m.Product{}, err
	}

	return product, nil
}

func (db *productRepository) GetProducts() ([]m.Product, error) {
	var products []m.Product

	rows, err := db.db.Query(conf.GetProductsQuery)
	if err != nil {
		return []m.Product{}, err
	}

	for rows.Next() {
		var product m.Product
		err := rows.Scan(&product.ProductID, &product.Name, &product.Unit, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return []m.Product{}, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (db *productRepository) UpdateProduct(product m.Product) (m.Product, error) {
	err := db.db.QueryRow(conf.UpdateProductQuery, product.Name, product.Unit, product.Price, product.CreatedAt, product.UpdatedAt, product.ProductID)
	if err != nil {
		return m.Product{}, fmt.Errorf("error updating service")
	}

	return product, nil
}

func (db *productRepository) DeleteProduct(id int) error {
	_, err := db.db.Exec(conf.DeleteProductQuery, id)
	if err != nil {
		return fmt.Errorf("error deleting service")
	}
	return nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db}
}
