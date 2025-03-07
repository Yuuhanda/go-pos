package model

import "time"

// Member represents the member table in the database
type Member struct {
	ID           int       `json:"id_member" db:"id_member"`
	Name         string    `json:"name_member" db:"name_member"`
	Phone        int       `json:"phone_member" db:"phone_member"`
	JoinDate     time.Time `json:"join_date" db:"join_date"`
	Token        string    `json:"token" db:"token"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	Points       int       `json:"member_point" db:"member_point(32)"`
}
