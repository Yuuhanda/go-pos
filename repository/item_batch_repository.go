package repository

import (
	"database/sql"
	"fmt"
	"go-pos/database"
	"go-pos/model"
)

// ItemBatchRepository handles database operations for item batches
type ItemBatchRepository struct{}

// NewItemBatchRepository creates a new ItemBatchRepository
func NewItemBatchRepository() *ItemBatchRepository {
	return &ItemBatchRepository{}
}

// CreateItemBatch inserts a new item batch into the database
func (r *ItemBatchRepository) CreateItemBatch(itemBatch *model.ItemBatch) (*model.ItemBatch, error) {
	query := `INSERT INTO item_batch (id_item, date_in, date_out, batch_qty) 
	          VALUES (?, ?, ?, ?)`
	          
	result, err := database.DB.Exec(query, 
		itemBatch.ItemID, 
		itemBatch.DateIn, 
		itemBatch.DateOut, 
		itemBatch.Qty)
		
	if err != nil {
		return nil, err
	}
	
	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	itemBatch.ID = int(lastID)
	return itemBatch, nil
}

// GetItemBatch retrieves an item batch by ID from the database
func (r *ItemBatchRepository) GetItemBatch(id int) (*model.ItemBatch, error) {
	itemBatch := &model.ItemBatch{}
	
	query := `SELECT id_batch, id_item, date_in, date_out, batch_qty 
	          FROM item_batch WHERE id_batch = ?`
	          
	err := database.DB.QueryRow(query, id).Scan(
		&itemBatch.ID,
		&itemBatch.ItemID,
		&itemBatch.DateIn,
		&itemBatch.DateOut,
		&itemBatch.Qty,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("item batch with ID %d not found", id)
		}
		return nil, err
	}
	
	return itemBatch, nil
}

// GetAllItemBatches retrieves all item batches from the database
func (r *ItemBatchRepository) GetAllItemBatches() ([]model.ItemBatch, error) {
	var itemBatches []model.ItemBatch
	
	query := `SELECT id_batch, id_item, date_in, date_out, batch_qty 
	          FROM item_batch ORDER BY date_in DESC`
	          
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var itemBatch model.ItemBatch
		err := rows.Scan(
			&itemBatch.ID,
			&itemBatch.ItemID,
			&itemBatch.DateIn,
			&itemBatch.DateOut,
			&itemBatch.Qty,
		)
		
		if err != nil {
			return nil, err
		}
		
		itemBatches = append(itemBatches, itemBatch)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return itemBatches, nil
}

// GetItemBatchesByItem retrieves all batches for a specific item
func (r *ItemBatchRepository) GetItemBatchesByItem(itemID int) ([]model.ItemBatch, error) {
	var itemBatches []model.ItemBatch
	
	query := `SELECT id_batch, id_item, date_in, date_out, batch_qty 
	          FROM item_batch 
	          WHERE id_item = ? 
	          ORDER BY date_in DESC`
	          
	rows, err := database.DB.Query(query, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var itemBatch model.ItemBatch
		err := rows.Scan(
			&itemBatch.ID,
			&itemBatch.ItemID,
			&itemBatch.DateIn,
			&itemBatch.DateOut,
			&itemBatch.Qty,
		)
		
		if err != nil {
			return nil, err
		}
		
		itemBatches = append(itemBatches, itemBatch)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return itemBatches, nil
}

// UpdateItemBatch updates an existing item batch in the database
func (r *ItemBatchRepository) UpdateItemBatch(itemBatch *model.ItemBatch) (*model.ItemBatch, error) {
	query := `UPDATE item_batch SET 
	          id_item = ?, 
	          date_in = ?, 
	          date_out = ?, 
	          batch_qty = ? 
	          WHERE id_batch = ?`
	          
	_, err := database.DB.Exec(query,
		itemBatch.ItemID,
		itemBatch.DateIn,
		itemBatch.DateOut,
		itemBatch.Qty,
		itemBatch.ID)

	if err != nil {
		return nil, err
	}

	return itemBatch, nil
}

// DeleteItemBatch deletes an item batch from the database
func (r *ItemBatchRepository) DeleteItemBatch(id int) error {
	query := `DELETE FROM item_batch WHERE id_batch = ?`
	
	_, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}
	
	return nil
}
