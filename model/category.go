package model

// Category represents the category table in the database
type Category struct {
	ID   int    `json:"id_category" db:"id_category"`
	Name string `json:"category_name" db:"category_name"`
}
