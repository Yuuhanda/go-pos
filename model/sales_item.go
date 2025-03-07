package model

// SalesItem represents the sales_item table in the database
type SalesItem struct {
	ID          int `json:"id_sales_item" db:"id_sales_item"`
	SalesID     int `json:"id_sales" db:"id_sales"`
	ItemID      int `json:"id_item" db:"id_item"`
	Qty         int `json:"qty" db:"qty"`
	TotalAmount int `json:"total_item_sales" db:"total_item_sales"`
	
	// Optional relation fields (not in database)
	Sales       *SalesBasket `json:"sales,omitempty" db:"-"`
	Item        *Item        `json:"item,omitempty" db:"-"`
}
