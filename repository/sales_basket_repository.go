package repository

import (
	"database/sql"
	"fmt"
	"go-pos/database"
	"go-pos/model"
)

// SalesBasketRepository handles database operations for sales baskets
type SalesBasketRepository struct{}

// NewSalesBasketRepository creates a new SalesBasketRepository
func NewSalesBasketRepository() *SalesBasketRepository {
	return &SalesBasketRepository{}
}

// CreateSalesBasketTx inserts a new sales basket as part of a transaction
func (r *SalesBasketRepository) CreateSalesBasketTx(tx *sql.Tx, basket *model.SalesBasket) (*model.SalesBasket, error) {
	query := `INSERT INTO sales_basket (id_user, id_member, sales_date, payment_method, total_amount) 
	          VALUES (?, ?, ?, ?, ?)`
	          
	result, err := tx.Exec(query, 
		basket.UserID, 
		basket.MemberID, 
		basket.SalesDate, 
		basket.PaymentMethod,
		basket.Total)
		
	if err != nil {
		return nil, err
	}
	
	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	basket.ID = int(lastID)
	return basket, nil
}

// GetSalesBasket retrieves a sales basket by ID from the database
func (r *SalesBasketRepository) GetSalesBasket(id int) (*model.SalesBasket, error) {
	basket := &model.SalesBasket{}
	
	query := `SELECT id_sales, id_user, id_member, sales_date, payment_method, total_amount 
	          FROM sales_basket WHERE id_sales = ?`
	          
	err := database.DB.QueryRow(query, id).Scan(
		&basket.ID,
		&basket.UserID,
		&basket.MemberID,
		&basket.SalesDate,
		&basket.PaymentMethod,
		&basket.Total,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("sales basket with ID %d not found", id)
		}
		return nil, err
	}
	
	return basket, nil
}

// GetAllSalesBaskets retrieves all sales baskets from the database
func (r *SalesBasketRepository) GetAllSalesBaskets() ([]model.SalesBasket, error) {
	var baskets []model.SalesBasket
	
	query := `SELECT id_sales, id_user, id_member, sales_date, payment_method, total_amount 
	          FROM sales_basket ORDER BY sales_date DESC`
	          
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var basket model.SalesBasket
		err := rows.Scan(
			&basket.ID,
			&basket.UserID,
			&basket.MemberID,
			&basket.SalesDate,
			&basket.PaymentMethod,
			&basket.Total,
		)
		
		if err != nil {
			return nil, err
		}
		
		baskets = append(baskets, basket)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return baskets, nil
}

// GetSalesBasketsByUser retrieves all sales baskets for a specific user
func (r *SalesBasketRepository) GetSalesBasketsByUser(userID int) ([]model.SalesBasket, error) {
	var baskets []model.SalesBasket
	
	query := `SELECT id_sales, id_user, id_member, sales_date, payment_method, total_amount 
	          FROM sales_basket 
	          WHERE id_user = ? 
	          ORDER BY sales_date DESC`
	          
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var basket model.SalesBasket
		err := rows.Scan(
			&basket.ID,
			&basket.UserID,
			&basket.MemberID,
			&basket.SalesDate,
			&basket.PaymentMethod,
			&basket.Total,
		)
		
		if err != nil {
			return nil, err
		}
		
		baskets = append(baskets, basket)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return baskets, nil
}

// GetSalesBasketsByMember retrieves all sales baskets for a specific member
func (r *SalesBasketRepository) GetSalesBasketsByMember(memberID int) ([]model.SalesBasket, error) {
	var baskets []model.SalesBasket
	
	query := `SELECT id_sales, id_user, id_member, sales_date, payment_method, total_amount 
	          FROM sales_basket 
	          WHERE id_member = ? 
	          ORDER BY sales_date DESC`
	          
	rows, err := database.DB.Query(query, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var basket model.SalesBasket
		err := rows.Scan(
			&basket.ID,
			&basket.UserID,
			&basket.MemberID,
			&basket.SalesDate,
			&basket.PaymentMethod,
			&basket.Total,
		)
		
		if err != nil {
			return nil, err
		}
		
		baskets = append(baskets, basket)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return baskets, nil
}

// GetSalesBasketsByUserAndMember retrieves all sales baskets for a specific user and member
func (r *SalesBasketRepository) GetSalesBasketsByUserAndMember(userID, memberID int) ([]model.SalesBasket, error) {
	var baskets []model.SalesBasket
	
	query := `SELECT id_sales, id_user, id_member, sales_date, payment_method, total_amount 
	          FROM sales_basket 
	          WHERE id_user = ? AND id_member = ? 
	          ORDER BY sales_date DESC`
	          
	rows, err := database.DB.Query(query, userID, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var basket model.SalesBasket
		err := rows.Scan(
			&basket.ID,
			&basket.UserID,
			&basket.MemberID,
			&basket.SalesDate,
			&basket.PaymentMethod,
			&basket.Total,
		)
		
		if err != nil {
			return nil, err
		}
		
		baskets = append(baskets, basket)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return baskets, nil
}

// UpdateSalesBasket updates an existing sales basket in the database
func (r *SalesBasketRepository) UpdateSalesBasket(basket *model.SalesBasket) (*model.SalesBasket, error) {
	query := `UPDATE sales_basket SET 
	          id_user = ?, 
	          id_member = ?, 
	          sales_date = ?, 
	          payment_method = ?, 
	          total_amount = ? 
	          WHERE id_sales = ?`
	          
	_, err := database.DB.Exec(query,
		basket.UserID,
		basket.MemberID,
		basket.SalesDate,
		basket.PaymentMethod,
		basket.Total,
		basket.ID)
		
	if err != nil {
		return nil, err
	}
	
	return basket, nil
}

// DeleteSalesBasketTx deletes a sales basket as part of a transaction
func (r *SalesBasketRepository) DeleteSalesBasketTx(tx *sql.Tx, id int) error {
	query := `DELETE FROM sales_basket WHERE id_sales = ?`
	
	_, err := tx.Exec(query, id)
	if err != nil {
		return err
	}
	
	return nil
}
