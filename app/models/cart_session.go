package models

import "time"

type SessionStatus string

var ACTIVE SessionStatus = "active"
var INACTIVE SessionStatus = "inactive"

type CartSession struct {
	Id        string        `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID    string        `json:"user_id" gorm:"not null;column:user_id;default:null"`
	Total     int           `json:"total" gorm:"column:total;default:null"`
	Status    SessionStatus `json:"status" gorm:"column:status;default:null"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
