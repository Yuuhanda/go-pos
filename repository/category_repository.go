package repository

import (
	"database/sql"
	"fmt"
	"go-pos/database"
	"go-pos/model"
)

// CategoryRepository handles database operations for categories
type CategoryRepository struct{}

// NewCategoryRepository creates a new CategoryRepository
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

// CreateCategory inserts a new category into the database
func (r *CategoryRepository) CreateCategory(category *model.Category) (*model.Category, error) {
	query := `INSERT INTO category (category_name) VALUES (?)`
	          
	result, err := database.DB.Exec(query, category.Name)
	if err != nil {
		return nil, err
	}
	
	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	category.ID = int(lastID)
	return category, nil
}

// GetCategory retrieves a category by ID from the database
func (r *CategoryRepository) GetCategory(id int) (*model.Category, error) {
	category := &model.Category{}
	
	query := `SELECT id_category, category_name FROM category WHERE id_category = ?`
	          
	err := database.DB.QueryRow(query, id).Scan(&category.ID, &category.Name)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("category with ID %d not found", id)
		}
		return nil, err
	}
	
	return category, nil
}

// GetAllCategories retrieves all categories from the database
func (r *CategoryRepository) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	
	query := `SELECT id_category, category_name FROM category ORDER BY category_name`
	          
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.ID, &category.Name)
		
		if err != nil {
			return nil, err
		}
		
		categories = append(categories, category)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return categories, nil
}

// UpdateCategory updates an existing category in the database
func (r *CategoryRepository) UpdateCategory(category *model.Category) (*model.Category, error) {
	query := `UPDATE category SET category_name = ? WHERE id_category = ?`
	          
	_, err := database.DB.Exec(query, category.Name, category.ID)
	if err != nil {
		return nil, err
	}
	
	return category, nil
}

// IsCategoryInUse checks if a category is being used by any items
func (r *CategoryRepository) IsCategoryInUse(id int) (bool, error) {
	var count int
	
	query := `SELECT COUNT(*) FROM item WHERE item_category = ?`
	
	err := database.DB.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	
	return count > 0, nil
}

// DeleteCategory deletes a category from the database
func (r *CategoryRepository) DeleteCategory(id int) error {
	query := `DELETE FROM category WHERE id_category = ?`
	
	_, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}
	
	return nil
}
