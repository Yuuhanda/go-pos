package repository

import (
	"go-pos/database"
	"go-pos/model"
	"database/sql"
	"fmt"
)

// UserRepository handles database operations for users
type UserRepository struct{}

// NewUserRepository creates a new UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	query := `INSERT INTO user (nik, name, address, phone, gender, password_hash, is_admin, token) 
		      VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
		      
	result, err := database.DB.Exec(query, 
		user.NIK, 
		user.Name, 
		user.Address, 
		user.Phone, 
		user.Gender, 
		user.PasswordHash,
		user.IsAdmin,
		user.Token)
		
	if err != nil {
		return nil, err
	}
	
	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	user.ID = int(lastID)
	return user, nil
}

// GetUser retrieves a user by ID from the database
func (r *UserRepository) GetUser(id int) (*model.User, error) {
	user := &model.User{}
	
	query := `SELECT id_user, nik, name, address, phone, gender, password_hash, is_admin, token 
	          FROM user WHERE id_user = ?`
	          
	err := database.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.NIK,
		&user.Name,
		&user.Address,
		&user.Phone,
		&user.Gender,
		&user.PasswordHash,
		&user.IsAdmin,
		&user.Token,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, err
	}
	
	return user, nil
}

