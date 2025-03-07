package model

// Gender defines the gender type
type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
)

// User represents the user table in the database
type User struct {
	ID           int    `json:"id" db:"id"`
	NIK          int    `json:"nik" db:"nik"`
	Name         string `json:"name" db:"name"`
	Address      string `json:"address" db:"address"`
	Phone        int    `json:"phone" db:"phone"`
	Gender       Gender `json:"gender" db:"gender"`
	IsAdmin      bool   `json:"admin" db:"admin"` // tinyint(32) converted to bool
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Token        string `json:"token" db:"token"`
}
