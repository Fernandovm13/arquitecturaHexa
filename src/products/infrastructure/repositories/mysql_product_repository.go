package repositories

import (
	"errors"
	"database/sql"
	"holamundo/src/core"
	"holamundo/src/products/domain/entities"
)

type MySQLProductRepository struct {
	db *sql.DB
}

func NewMySQLProductRepository() *MySQLProductRepository {
	return &MySQLProductRepository{db: core.GetDB()}
}

func (repo *MySQLProductRepository) Save(product *entities.Product) error {
	query := "INSERT INTO products (name, price, category_id) VALUES (?, ?, ?)"
	_, err := repo.db.Exec(query, product.Name, product.Price, product.CategoryID)
	return err
}

func (repo *MySQLProductRepository) GetAll() ([]entities.Product, error) {
	query := "SELECT id, name, price, category_id FROM products"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *MySQLProductRepository) Update(product *entities.Product) error {
	query := "UPDATE products SET name=?, price=?, category_id=? WHERE id=?"
	_, err := repo.db.Exec(query, product.Name, product.Price, product.CategoryID, product.ID)
	return err
}

func (repo *MySQLProductRepository) Delete(id int32) error {
	query := "DELETE FROM products WHERE id=?"
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *MySQLProductRepository) GetByID(id int32) (*entities.Product, error) {
    query := "SELECT id, name, price, category_id FROM products WHERE id = ?"
    row := repo.db.QueryRow(query, id)

    var product entities.Product
    err := row.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("producto no encontrado")
        }
        return nil, err
    }

    return &product, nil
}