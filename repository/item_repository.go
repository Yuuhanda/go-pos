package repository

import (
    "database/sql"
    "fmt"
    "go-pos/database"
    "go-pos/model"
)

// ItemRepository handles database operations for items
type ItemRepository struct{}

// NewItemRepository creates a new ItemRepository
func NewItemRepository() *ItemRepository {
    return &ItemRepository{}
}

// CreateItem inserts a new item into the database
func (r *ItemRepository) CreateItem(item *model.Item) (*model.Item, error) {
    query := `INSERT INTO item (item_category, item_name, item_price) 
              VALUES (?, ?, ?)`
              
    result, err := database.DB.Exec(query, 
        item.CategoryID, 
        item.Name, 
        item.Price)
        
    if err != nil {
        return nil, err
    }
    
    // Get the last inserted ID
    lastID, err := result.LastInsertId()
    if err != nil {
        return nil, err
    }
    
    item.ID = int(lastID)
    return item, nil
}

// GetItem retrieves an item by ID from the database
func (r *ItemRepository) GetItem(id int) (*model.Item, error) {
    item := &model.Item{}
    
    query := "SELECT id_item, item_category, item_name, item_price FROM item WHERE id_item = ?"
    err := database.DB.QueryRow(query, id).Scan(&item.ID, &item.CategoryID, &item.Name, &item.Price)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("item with ID %d not found", id)
        }
        return nil, err
    }
    
    return item, nil
}

// GetAllItems retrieves all items from the database
func (r *ItemRepository) GetAllItems() ([]model.Item, error) {
    var items []model.Item
    
    query := `SELECT id_item, item_category, item_name, item_price 
              FROM item ORDER BY item_name`
              
    rows, err := database.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    for rows.Next() {
        var item model.Item
        err := rows.Scan(
            &item.ID,
            &item.CategoryID,
            &item.Name,
            &item.Price,
        )
        
        if err != nil {
            return nil, err
        }
        
        items = append(items, item)
    }
    
    if err = rows.Err(); err != nil {
        return nil, err
    }
    
    return items, nil
}

// GetItemsByCategory retrieves all items in a specific category
func (r *ItemRepository) GetItemsByCategory(categoryID int) ([]model.Item, error) {
    var items []model.Item
    
    query := `SELECT id_item, item_category, item_name, item_price 
              FROM item WHERE item_category = ? 
              ORDER BY item_name`
              
    rows, err := database.DB.Query(query, categoryID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    for rows.Next() {
        var item model.Item
        err := rows.Scan(
            &item.ID,
            &item.CategoryID,
            &item.Name,
            &item.Price,
        )
        
        if err != nil {
            return nil, err
        }
        
        items = append(items, item)
    }
    
    if err = rows.Err(); err != nil {
        return nil, err
    }
    
    return items, nil
}

// UpdateItem updates an existing item in the database
func (r *ItemRepository) UpdateItem(item *model.Item) (*model.Item, error) {
    query := `UPDATE item SET 
              item_category = ?, 
              item_name = ?, 
              item_price = ? 
              WHERE id_item = ?`
              
    _, err := database.DB.Exec(query,
        item.CategoryID,
        item.Name,
        item.Price,
        item.ID)
        
    if err != nil {
        return nil, err
    }
    
    return item, nil
}

// DeleteItem deletes an item from the database
func (r *ItemRepository) DeleteItem(id int) error {
    query := `DELETE FROM item WHERE id_item = ?`
    
    _, err := database.DB.Exec(query, id)
    if err != nil {
        return err
    }
    
    return nil
}
