package model

// Item represents the item table in the database
type Item struct {
	ID         int    `json:"id_item" db:"id_item"`
	CategoryID int    `json:"item_category" db:"item_category"`
	Name       string `json:"item_name" db:"item_name"`
	Price      int    `json:"item_price" db:"item_price"`
	
	// Optional relation field (not in database)
	Category   *Category `json:"category,omitempty" db:"-"`
}
