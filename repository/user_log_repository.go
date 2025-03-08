package repository

import (
	"database/sql"
	"fmt"
	"go-pos/database"
	"go-pos/model"
)

// UserLogRepository handles database operations for user logs
type UserLogRepository struct{}

// NewUserLogRepository creates a new UserLogRepository
func NewUserLogRepository() *UserLogRepository {
	return &UserLogRepository{}
}

// CreateUserLog inserts a new user log into the database
func (r *UserLogRepository) CreateUserLog(userLog *model.UserLog) (*model.UserLog, error) {
	query := `INSERT INTO user_log (id_user, date, ip, platform_browser, action) 
		      VALUES (?, ?, ?, ?, ?)`
		      
	result, err := database.DB.Exec(query, 
		userLog.UserID, 
		userLog.Date,
		userLog.IP,
		userLog.PlatformBrowser,
		userLog.Action)
		
	if err != nil {
		return nil, err
	}
	
	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	userLog.ID = int(lastID)
	return userLog, nil
}

// GetUserLog retrieves a user log by ID from the database
func (r *UserLogRepository) GetUserLog(id int) (*model.UserLog, error) {
	userLog := &model.UserLog{}
	
	query := `SELECT id_log, id_user, date, ip, platform_browser, action
	          FROM user_log WHERE id_log = ?`
	          
	err := database.DB.QueryRow(query, id).Scan(
		&userLog.ID,
		&userLog.UserID,
		&userLog.Date,
		&userLog.IP,
		&userLog.PlatformBrowser,
		&userLog.Action,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user log with ID %d not found", id)
		}
		return nil, err
	}
	
	return userLog, nil
}

// GetAllUserLogs retrieves all user logs from the

// GetUserLogsByUserID retrieves all user logs for a specific user
func (r *UserLogRepository) GetUserLogsByUserID(userID int) ([]model.UserLog, error) {
    var userLogs []model.UserLog
    
    query := `SELECT id_log, id_user, date, ip, platform_browser 
              FROM user_log 
              WHERE id_user = ?
              ORDER BY date DESC`
              
    rows, err := database.DB.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    for rows.Next() {
        var userLog model.UserLog
        err := rows.Scan(
            &userLog.ID,
            &userLog.UserID,
            &userLog.Date,
            &userLog.IP,
            &userLog.PlatformBrowser,
        )
        
        if err != nil {
            return nil, err
        }
        
        userLogs = append(userLogs, userLog)
    }
    
    if err = rows.Err(); err != nil {
        return nil, err
    }
    
    return userLogs, nil
}

// GetAllUserLogs retrieves all user logs from the database
func (r *UserLogRepository) GetAllUserLogs() ([]model.UserLog, error) {
    var userLogs []model.UserLog
    
    query := `SELECT id_log, id_user, date, ip, platform_browser 
              FROM user_log 
              ORDER BY date DESC`
              
    rows, err := database.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    for rows.Next() {
        var userLog model.UserLog
        err := rows.Scan(
            &userLog.ID,
            &userLog.UserID,
            &userLog.Date,
            &userLog.IP,
            &userLog.PlatformBrowser,
        )
        
        if err != nil {
            return nil, err
        }
        
        userLogs = append(userLogs, userLog)
    }
    
    if err = rows.Err(); err != nil {
        return nil, err
    }
    
    return userLogs, nil
}

// UpdateUserLog updates an existing user log in the database
func (r *UserLogRepository) UpdateUserLog(userLog *model.UserLog) (*model.UserLog, error) {
    query := `UPDATE user_log SET 
              id_user = ?, 
              date = ?, 
              ip = ?, 
              platform_browser = ? 
              WHERE id_log = ?`
              
    _, err := database.DB.Exec(query,
        userLog.UserID,
        userLog.Date,
        userLog.IP,
        userLog.PlatformBrowser,
        userLog.ID)
        
    if err != nil {
        return nil, err
    }
    
    return userLog, nil
}

// DeleteUserLog deletes a user log from the database
func (r *UserLogRepository) DeleteUserLog(id int) error {
    query := `DELETE FROM user_log WHERE id_log = ?`
    
    _, err := database.DB.Exec(query, id)
    if err != nil {
        return err
    }
    
    return nil
}

