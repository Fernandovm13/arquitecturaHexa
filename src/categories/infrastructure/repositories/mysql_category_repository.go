package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"holamundo/src/categories/domain/entities"
	"holamundo/src/core"
)

type MySQLCategoryRepository struct {
	db *sql.DB
}

func NewMySQLCategoryRepository() *MySQLCategoryRepository {
	return &MySQLCategoryRepository{db: core.GetDB()}
}

func (repo *MySQLCategoryRepository) Save(category *entities.Category) error {
	query := "INSERT INTO categories (name, secret) VALUES (?, ?)"
	_, err := repo.db.Exec(query, category.Name, category.Secret)
	if err != nil {
		return fmt.Errorf("error al insertar la categoría: %v", err)
	}
	return nil
}

func (repo *MySQLCategoryRepository) GetAll() ([]entities.Category, error) {
	query := "SELECT id, name, secret FROM categories"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todas las categorías: %v", err)
	}
	defer rows.Close()

	var categories []entities.Category
	for rows.Next() {
		var category entities.Category
		var secret sql.NullString
		if err := rows.Scan(&category.ID, &category.Name, &secret); err != nil {
			return nil, fmt.Errorf("error al leer fila: %v", err)
		}
		category.Secret = secret
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar filas: %v", err)
	}

	return categories, nil
}

func (repo *MySQLCategoryRepository) GetByID(id int32) (*entities.Category, error) {
	query := "SELECT id, name, secret FROM categories WHERE id = ?"
	row := repo.db.QueryRow(query, id)

	var category entities.Category
	var secret sql.NullString
	err := row.Scan(&category.ID, &category.Name, &secret)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("categoría no encontrada")
		}
		return nil, err
	}
	category.Secret = secret

	return &category, nil
}

func (repo *MySQLCategoryRepository) Update(category *entities.Category) error {
	query := "UPDATE categories SET name = ?, secret = ? WHERE id = ?"
	_, err := repo.db.Exec(query, category.Name, category.Secret, category.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar la categoría: %v", err)
	}
	return nil
}

func (repo *MySQLCategoryRepository) Delete(id int32) error {
	query := "DELETE FROM categories WHERE id = ?"
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar la categoría: %v", err)
	}
	return nil
}
