package repository

import (
	"database/sql"
	"fmt"
	"go-pos/database"
	"go-pos/model"
)

// SalesItemRepository handles database operations for sales items
type SalesItemRepository struct{}

// NewSalesItemRepository creates a new SalesItemRepository
func NewSalesItemRepository() *SalesItemRepository {
	return &SalesItemRepository{}
}

// CreateSalesItem inserts a new sales item into the database
func (r *SalesItemRepository) CreateSalesItem(item *model.SalesItem) (*model.SalesItem, error) {
	query := `INSERT INTO sales_item (id_sales, id_item, qty, total_item_sales) 
	          VALUES (?, ?, ?, ?)`
	          
	result, err := database.DB.Exec(query, 
		item.SalesID, 
		item.ItemID, 
		item.Qty, 
		item.TotalAmount)
		
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

// CreateSalesItemTx inserts a new sales item as part of a transaction
func (r *SalesItemRepository) CreateSalesItemTx(tx *sql.Tx, item *model.SalesItem) (*model.SalesItem, error) {
	query := `INSERT INTO sales_item (id_sales, id_item, qty, total_item_sales) 
	          VALUES (?, ?, ?, ?)`
	          
	result, err := tx.Exec(query, 
		item.SalesID, 
		item.ItemID, 
		item.Qty, 
		item.TotalAmount)
		
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

// GetSalesItem retrieves a sales item by ID from the database
func (r *SalesItemRepository) GetSalesItem(id int) (*model.SalesItem, error) {
	salesItem := &model.SalesItem{}
	
	query := `SELECT id_sales_item, id_sales, id_item, qty, total_item_sales 
	          FROM sales_item WHERE id_sales_item = ?`
	          
	err := database.DB.QueryRow(query, id).Scan(
		&salesItem.ID,
		&salesItem.SalesID,
		&salesItem.ItemID,
		&salesItem.Qty,
		&salesItem.TotalAmount,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("sales item with ID %d not found", id)
		}
		return nil, err
	}
	
	return salesItem, nil
}

// GetSalesItemsBySales retrieves all sales items for a specific sales basket
func (r *SalesItemRepository) GetSalesItemsBySales(salesID int) ([]model.SalesItem, error) {
	var salesItems []model.SalesItem
	
	query := `SELECT id_sales_item, id_sales, id_item, qty, total_item_sales 
	          FROM sales_item 
	          WHERE id_sales = ?`
	          
	rows, err := database.DB.Query(query, salesID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var salesItem model.SalesItem
		err := rows.Scan(
			&salesItem.ID,
			&salesItem.SalesID,
			&salesItem.ItemID,
			&salesItem.Qty,
			&salesItem.TotalAmount,
		)
		
		if err != nil {
			return nil, err
		}
		
		salesItems = append(salesItems, salesItem)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return salesItems, nil
}

// GetAllSalesItems retrieves all sales items from the database
func (r *SalesItemRepository) GetAllSalesItems() ([]model.SalesItem, error) {
	var salesItems []model.SalesItem
	
	query := `SELECT id_sales_item, id_sales, id_item, qty, total_item_sales 
	          FROM sales_item`
	          
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var salesItem model.SalesItem
		err := rows.Scan(
			&salesItem.ID,
			&salesItem.SalesID,
			&salesItem.ItemID,
			&salesItem.Qty,
			&salesItem.TotalAmount,
		)
		
		if err != nil {
			return nil, err
		}
		
		salesItems = append(salesItems, salesItem)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return salesItems, nil
}

// UpdateSalesItem updates an existing sales item in the database
func (r *SalesItemRepository) UpdateSalesItem(salesItem *model.SalesItem) (*model.SalesItem, error) {
	query := `UPDATE sales_item SET 
	          id_sales = ?, 
	          id_item = ?, 
	          qty = ?, 
	          total_item_sales = ? 
	          WHERE id_sales_item = ?`
	          
	_, err := database.DB.Exec(query,
		salesItem.SalesID,
		salesItem.ItemID,
		salesItem.Qty,
		salesItem.TotalAmount,
		salesItem.ID)
		
	if err != nil {
		return nil, err
	}
	
	return salesItem, nil
}

// DeleteSalesItem deletes a sales item from the database
func (r *SalesItemRepository) DeleteSalesItem(id int) error {
	query := `DELETE FROM sales_item WHERE id_sales_item = ?`
	
	_, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}
	
	return nil
}

// DeleteSalesItemsBySalesTx deletes all sales items for a specific sales basket as part of a transaction
func (r *SalesItemRepository) DeleteSalesItemsBySalesTx(tx *sql.Tx, salesID int) error {
	query := `DELETE FROM sales_item WHERE id_sales = ?`
	
	_, err := tx.Exec(query, salesID)
	if err != nil {
		return err
	}
	
	return nil
}
