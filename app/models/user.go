package models

import "time"

type Role string

var ADMIN Role = "admin"
var USER Role = "user"

type User struct {
	ID             string    `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Username       string    `json:"username" gorm:"not null;column:username;default:null"`
	PasswordDigest string    `json:"-" gorm:"not null;column:password_digest;default:null"`
	Role           Role      `json:"role" gorm:"not null;column:role;default:null"`
	Active         bool      `json:"active" gorm:"not null;column:active;default:null"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
