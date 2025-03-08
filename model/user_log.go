package model

import "time"

// UserLog represents the user_log table in the database
type UserLog struct {
	ID              int       `json:"id_log_user" db:"id_log_user"`
	UserID          int       `json:"id_user" db:"id_user"`
	Date            time.Time `json:"date" db:"date"`
	IP              string    `json:"ip" db:"ip"`
	PlatformBrowser string    `json:"platform_browser" db:"platform_browser"`
	Action          string    `json:"action" db:"action"`
	
	// Optional relation field (not in database)
	User            *User     `json:"user,omitempty" db:"-"`
}
