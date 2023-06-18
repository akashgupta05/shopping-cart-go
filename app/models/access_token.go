package models

import "time"

type AccessToken struct {
	ID        string    `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID    string    `json:"user_id" gorm:"not null;column:user_id;default:null"`
	Token     string    `json:"token" gorm:"column:token;default:null"`
	Active    bool      `json:"active" gorm:"column:active;default:null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
