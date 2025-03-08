package repository

import (
	"database/sql"
	"fmt"
	"go-pos/database"
	"go-pos/model"
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

// GetAllUsers retrieves all users from the database
func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	
	query := `SELECT id_user, nik, name, address, phone, gender, password_hash, is_admin, token 
	          FROM user`
	          
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var user model.User
		err := rows.Scan(
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
			return nil, err
		}
		
		users = append(users, user)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return users, nil
}

// UpdateUser updates an existing user in the database
func (r *UserRepository) UpdateUser(user *model.User) (*model.User, error) {
	query := `UPDATE user SET 
	          nik = ?, 
	          name = ?, 
	          address = ?, 
	          phone = ?, 
	          gender = ?, 
	          password_hash = ?, 
	          is_admin = ?, 
	          token = ? 
	          WHERE id_user = ?`
	          
	_, err := database.DB.Exec(query,
		user.NIK,
		user.Name,
		user.Address,
		user.Phone,
		user.Gender,
		user.PasswordHash,
		user.IsAdmin,
		user.Token,
		user.ID)
		
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

// DeleteUser deletes a user from the database
func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM user WHERE id_user = ?`
	
	_, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}
	
	return nil
}
