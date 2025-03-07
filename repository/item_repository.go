package repository

import (
    "go-pos/database"
    "go-pos/model"
)

// ItemRepository provides database operations for items
type ItemRepository struct{}

// GetItem retrieves an item by ID
func (r *ItemRepository) GetItem(id int) (*model.Item, error) {
    item := &model.Item{}
    
    query := "SELECT id_item, item_category, item_name, item_price FROM item WHERE id_item = ?"
    err := database.DB.QueryRow(query, id).Scan(&item.ID, &item.CategoryID, &item.Name, &item.Price)
    
    if err != nil {
        return nil, err
    }
    
    return item, nil
}

// Additional repository methods for CRUD operations
