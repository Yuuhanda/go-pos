package repository

import (
	"database/sql"
	"fmt"
	"go-pos/database"
	"go-pos/model"
)

// MemberRepository handles database operations for members
type MemberRepository struct{}

// NewMemberRepository creates a new MemberRepository
func NewMemberRepository() *MemberRepository {
	return &MemberRepository{}
}

// CreateMember inserts a new member into the database
func (r *MemberRepository) CreateMember(member *model.Member) (*model.Member, error) {
	query := `INSERT INTO member (member_name, member_phone, join_date, member_points) 
	          VALUES (?, ?, ?, ?)`
	          
	result, err := database.DB.Exec(query, 
		member.Name,
		member.Phone,
		member.JoinDate,
		member.Points)
		
	if err != nil {
		return nil, err
	}
	
	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	member.ID = int(lastID)
	return member, nil
}

// GetMember retrieves a member by ID from the database
func (r *MemberRepository) GetMember(id int) (*model.Member, error) {
	member := &model.Member{}
	
	query := `SELECT id_member, member_name, member_phone, join_date, member_points 
	          FROM member WHERE id_member = ?`
	          
	err := database.DB.QueryRow(query, id).Scan(
		&member.ID,
		&member.Name,
		&member.Phone,
		&member.JoinDate,
		&member.Points,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("member with ID %d not found", id)
		}
		return nil, err
	}
	
	return member, nil
}

// GetAllMembers retrieves all members from the database
func (r *MemberRepository) GetAllMembers() ([]model.Member, error) {
	var members []model.Member
	
	query := `SELECT id_member, member_name, member_phone, join_date, member_points 
	          FROM member ORDER BY member_name`
	          
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var member model.Member
		err := rows.Scan(
			&member.ID,
			&member.Name,
			&member.Phone,
			&member.JoinDate,
			&member.Points,
		)
		
		if err != nil {
			return nil, err
		}
		
		members = append(members, member)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return members, nil
}

// GetMembersByPhone retrieves members with a specific phone number
func (r *MemberRepository) GetMembersByPhone(phone int) ([]model.Member, error) {
	var members []model.Member
	
	query := `SELECT id_member, member_name, member_phone, join_date, member_points 
	          FROM member WHERE member_phone = ?`
	          
	rows, err := database.DB.Query(query, phone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var member model.Member
		err := rows.Scan(
			&member.ID,
			&member.Name,
			&member.Phone,
			&member.JoinDate,
			&member.Points,
		)
		
		if err != nil {
			return nil, err
		}
		
		members = append(members, member)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return members, nil
}

// UpdateMember updates an existing member in the database
func (r *MemberRepository) UpdateMember(member *model.Member) (*model.Member, error) {
	query := `UPDATE member SET 
	          member_name = ?, 
	          member_phone = ?, 
	          member_points = ? 
	          WHERE id_member = ?`
	          
	_, err := database.DB.Exec(query,
		member.Name,
		member.Phone,
		member.Points,
		member.ID)
		
	if err != nil {
		return nil, err
	}
	
	return member, nil
}

// MemberHasSales checks if a member has any sales records
func (r *MemberRepository) MemberHasSales(id int) (bool, error) {
	var count int
	
	query := `SELECT COUNT(*) FROM sales_basket WHERE id_member = ?`
	
	err := database.DB.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	
	return count > 0, nil
}

// DeleteMember deletes a member from the database
func (r *MemberRepository) DeleteMember(id int) error {
	query := `DELETE FROM member WHERE id_member = ?`
	
	_, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}
	
	return nil
}
