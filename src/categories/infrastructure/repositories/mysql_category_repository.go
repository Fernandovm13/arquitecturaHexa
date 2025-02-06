package repositories

import (
	"fmt"
	"errors"
	"database/sql"
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
	query := "INSERT INTO categories (name) VALUES (?)"
	_, err := repo.db.Exec(query, category.Name)
	if err != nil {
		fmt.Printf("Error al insertar la categoría: %v\n", err)
	}
	return err
}


func (repo *MySQLCategoryRepository) GetAll() ([]entities.Category, error) {
	query := "SELECT id, name FROM categories"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]entities.Category, 0)

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}


func (repo *MySQLCategoryRepository) Update(category *entities.Category) error {
	query := "UPDATE categories SET name=? WHERE id=?"
	_, err := repo.db.Exec(query, category.Name, category.ID)
	return err
}

func (repo *MySQLCategoryRepository) Delete(id int32) error {
	query := "DELETE FROM categories WHERE id=?"
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *MySQLCategoryRepository) GetByID(id int32) (*entities.Category, error) {
    query := "SELECT id, name FROM categories WHERE id = ?"
    row := repo.db.QueryRow(query, id)

    var category entities.Category
    err := row.Scan(&category.ID, &category.Name)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("categoría no encontrada")
        }
        return nil, err
    }

    return &category, nil
}
