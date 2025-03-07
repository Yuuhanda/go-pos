package model

import "time"

// UserMember represents the user_member table in the database
type UserMember struct {
	ID              int       `json:"id_log_member" db:"id_log_member"`
	MemberID        int       `json:"id_member" db:"id_member"`
	Date            time.Time `json:"date" db:"date"`
	IP              string    `json:"ip" db:"ip"`
	PlatformBrowser string    `json:"platform_browser" db:"platform_browser"`
	
	// Optional relation field (not in database)
	Member          *Member   `json:"member,omitempty" db:"-"`
}
