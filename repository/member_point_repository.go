package repository

import (
	"database/sql"
	"fmt"
	"go-pos/database"
	"go-pos/model"
	"time"
)

// MemberPointRepository handles database operations for member points
type MemberPointRepository struct{}

// NewMemberPointRepository creates a new MemberPointRepository
func NewMemberPointRepository() *MemberPointRepository {
	return &MemberPointRepository{}
}

// CreateMemberPoint inserts a new member point transaction into the database
func (r *MemberPointRepository) CreateMemberPoint(memberPoint *model.MemberPoint) (*model.MemberPoint, error) {
	query := `INSERT INTO member_point (id_member, type, points, transaction_date) 
	          VALUES (?, ?, ?, ?)`
	          
	result, err := database.DB.Exec(query, 
		memberPoint.MemberID,
		memberPoint.Type,
		memberPoint.Points,
		time.Now())
		
	if err != nil {
		return nil, err
	}
	
	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	memberPoint.ID = int(lastID)
	return memberPoint, nil
}

// CreateMemberPointTx inserts a new member point transaction as part of a transaction
func (r *MemberPointRepository) CreateMemberPointTx(tx *sql.Tx, memberPoint *model.MemberPoint) (*model.MemberPoint, error) {
	query := `INSERT INTO member_point (id_member, type, points, transaction_date) 
	          VALUES (?, ?, ?, ?)`
	          
	result, err := tx.Exec(query, 
		memberPoint.MemberID,
		memberPoint.Type,
		memberPoint.Points,
		time.Now())
		
	if err != nil {
		return nil, err
	}
	
	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	memberPoint.ID = int(lastID)
	return memberPoint, nil
}

// GetMemberPoint retrieves a member point transaction by ID
func (r *MemberPointRepository) GetMemberPoint(id int) (*model.MemberPoint, error) {
	memberPoint := &model.MemberPoint{}
	
	query := `SELECT id_point, id_member, type, points, transaction_date 
	          FROM member_point WHERE id_point = ?`
	          
	var transactionDate time.Time
	err := database.DB.QueryRow(query, id).Scan(
		&memberPoint.ID,
		&memberPoint.MemberID,
		&memberPoint.Type,
		&memberPoint.Points,
		&transactionDate,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("member point transaction with ID %d not found", id)
		}
		return nil, err
	}
	
	return memberPoint, nil
}

// GetAllMemberPoints retrieves all member point transactions
func (r *MemberPointRepository) GetAllMemberPoints() ([]model.MemberPoint, error) {
	var memberPoints []model.MemberPoint
	
	query := `SELECT id_point, id_member, type, points, transaction_date 
	          FROM member_point ORDER BY transaction_date DESC`
	          
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var memberPoint model.MemberPoint
		var transactionDate time.Time
		err := rows.Scan(
			&memberPoint.ID,
			&memberPoint.MemberID,
			&memberPoint.Type,
			&memberPoint.Points,
			&transactionDate,
		)
		
		if err != nil {
			return nil, err
		}
		
		memberPoints = append(memberPoints, memberPoint)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return memberPoints, nil
}

// GetMemberPointsByMember retrieves all member point transactions for a specific member
func (r *MemberPointRepository) GetMemberPointsByMember(memberID int) ([]model.MemberPoint, error) {
	var memberPoints []model.MemberPoint
	
	query := `SELECT id_point, id_member, type, points, transaction_date 
	          FROM member_point 
	          WHERE id_member = ? 
	          ORDER BY transaction_date DESC`
	          
	rows, err := database.DB.Query(query, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var memberPoint model.MemberPoint
		var transactionDate time.Time
		err := rows.Scan(
			&memberPoint.ID,
			&memberPoint.MemberID,
			&memberPoint.Type,
			&memberPoint.Points,
			&transactionDate,
		)
		
		if err != nil {
			return nil, err
		}
		
		memberPoints = append(memberPoints, memberPoint)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return memberPoints, nil
}

// GetMemberPointsByType retrieves all member point transactions of a specific type
func (r *MemberPointRepository) GetMemberPointsByType(pointType model.PointType) ([]model.MemberPoint, error) {
	var memberPoints []model.MemberPoint
	
	query := `SELECT id_point, id_member, type, points, transaction_date 
	          FROM member_point 
	          WHERE type = ? 
	          ORDER BY transaction_date DESC`
	          
	rows, err := database.DB.Query(query, pointType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var memberPoint model.MemberPoint
		var transactionDate time.Time
		err := rows.Scan(
			&memberPoint.ID,
			&memberPoint.MemberID,
			&memberPoint.Type,
			&memberPoint.Points,
			&transactionDate,
		)
		
		if err != nil {
			return nil, err
		}
		
		memberPoints = append(memberPoints, memberPoint)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return memberPoints, nil
}

// GetMemberPointsByMemberAndType retrieves all member point transactions for a specific member and type
func (r *MemberPointRepository) GetMemberPointsByMemberAndType(memberID int, pointType model.PointType) ([]model.MemberPoint, error) {
	var memberPoints []model.MemberPoint
	
	query := `SELECT id_point, id_member, type, points, transaction_date 
	          FROM member_point 
	          WHERE id_member = ? AND type = ? 
	          ORDER BY transaction_date DESC`
	          
	rows, err := database.DB.Query(query, memberID, pointType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var memberPoint model.MemberPoint
		var transactionDate time.Time
		err := rows.Scan(
			&memberPoint.ID,
			&memberPoint.MemberID,
			&memberPoint.Type,
			&memberPoint.Points,
			&transactionDate,
		)
		
		if err != nil {
			return nil, err
		}
		
		memberPoints = append(memberPoints, memberPoint)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return memberPoints, nil
}

// UpdateMemberPointTx updates a member point transaction as part of a transaction
func (r *MemberPointRepository) UpdateMemberPointTx(tx *sql.Tx, memberPoint *model.MemberPoint) (*model.MemberPoint, error) {
	query := `UPDATE member_point SET 
	          id_member = ?, 
	          type = ?, 
	          points = ? 
	          WHERE id_point = ?`
	          
	_, err := tx.Exec(query,
		memberPoint.MemberID,
		memberPoint.Type,
		memberPoint.Points,
		memberPoint.ID)
		
	if err != nil {
		return nil, err
	}
	
	return memberPoint, nil
}

// DeleteMemberPointTx deletes a member point transaction as part of a transaction
func (r *MemberPointRepository) DeleteMemberPointTx(tx *sql.Tx, id int) error {
	query := `DELETE FROM member_point WHERE id_point = ?`
	
	_, err := tx.Exec(query, id)
	if err != nil {
		return err
	}
	
	return nil
}