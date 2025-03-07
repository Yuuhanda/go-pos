package model

import "time"

// ItemBatch represents the item_batch table in the database
type ItemBatch struct {
	ID      int       `json:"id_batch" db:"id_batch"`
	ItemID  int       `json:"id_item" db:"id_item"`
	DateIn  time.Time `json:"date_in" db:"date_in"`
	DateOut time.Time `json:"date_out" db:"date_out"`
	Qty     int       `json:"batch_qty" db:"batch_qty"`
	
	// Optional relation field (not in database)
	Item    *Item     `json:"item,omitempty" db:"-"`
}
