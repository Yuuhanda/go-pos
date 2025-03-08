package repository

import (
	"database/sql"
	"fmt"
	"go-pos/database"
	"go-pos/model"
)

// UserMemberRepository handles database operations for user members
type UserMemberRepository struct{}

// NewUserMemberRepository creates a new UserMemberRepository
func NewUserMemberRepository() *UserMemberRepository {
	return &UserMemberRepository{}
}

// CreateUserMember inserts a new user member log into the database
func (r *UserMemberRepository) CreateUserMember(userMember *model.UserMember) (*model.UserMember, error) {
	query := `INSERT INTO user_member (id_member, date, ip, platform_browser) 
	          VALUES (?, ?, ?, ?)`
	          
	result, err := database.DB.Exec(query, 
		userMember.MemberID,
		userMember.Date,
		userMember.IP,
		userMember.PlatformBrowser)
		
	if err != nil {
		return nil, err
	}
	
	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	userMember.ID = int(lastID)
	return userMember, nil
}

// GetUserMember retrieves a user member log by ID
func (r *UserMemberRepository) GetUserMember(id int) (*model.UserMember, error) {
	userMember := &model.UserMember{}
	
	query := `SELECT id_user_member, id_member, date, ip, platform_browser 
	          FROM user_member WHERE id_user_member = ?`
	          
	err := database.DB.QueryRow(query, id).Scan(
		&userMember.ID,
		&userMember.MemberID,
		&userMember.Date,
		&userMember.IP,
		&userMember.PlatformBrowser,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user member log with ID %d not found", id)
		}
		return nil, err
	}
	
	return userMember, nil
}

// GetAllUserMembers retrieves all user member logs
func (r *UserMemberRepository) GetAllUserMembers() ([]model.UserMember, error) {
	var userMembers []model.UserMember
	
	query := `SELECT id_user_member, id_member, date, ip, platform_browser 
	          FROM user_member ORDER BY date DESC`
	          
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var userMember model.UserMember
		err := rows.Scan(
			&userMember.ID,
			&userMember.MemberID,
			&userMember.Date,
			&userMember.IP,
			&userMember.PlatformBrowser,
		)
		
		if err != nil {
			return nil, err
		}
		
		userMembers = append(userMembers, userMember)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return userMembers, nil
}

// GetUserMembersByMember retrieves all user member logs for a specific member
func (r *UserMemberRepository) GetUserMembersByMember(memberID int) ([]model.UserMember, error) {
	var userMembers []model.UserMember
	
	query := `SELECT id_user_member, id_member, date, ip, platform_browser 
	          FROM user_member 
	          WHERE id_member = ? 
	          ORDER BY date DESC`
	          
	rows, err := database.DB.Query(query, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var userMember model.UserMember
		err := rows.Scan(
			&userMember.ID,
			&userMember.MemberID,
			&userMember.Date,
			&userMember.IP,
			&userMember.PlatformBrowser,
		)
		
		if err != nil {
			return nil, err
		}
		
		userMembers = append(userMembers, userMember)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return userMembers, nil
}

// UpdateUserMember updates an existing user member log
func (r *UserMemberRepository) UpdateUserMember(userMember *model.UserMember) (*model.UserMember, error) {
	query := `UPDATE user_member SET 
	          id_member = ?, 
	          date = ?, 
	          ip = ?, 
	          platform_browser = ? 
	          WHERE id_user_member = ?`
	          
	_, err := database.DB.Exec(query,
		userMember.MemberID,
		userMember.Date,
		userMember.IP,
		userMember.PlatformBrowser,
		userMember.ID)
		
	if err != nil {
		return nil, err
	}
	
	return userMember, nil
}

// DeleteUserMember deletes a user member log
func (r *UserMemberRepository) DeleteUserMember(id int) error {
	query := `DELETE FROM user_member WHERE id_user_member = ?`
	
	_, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}
	
	return nil
}
