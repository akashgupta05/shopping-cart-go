package models

import "time"

type CartItem struct {
	Id        string    `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ItemID    string    `json:"item_id" gorm:"not null;column:item_id;default:null"`
	SessionID string    `json:"session_id" gorm:"not null;column:session_id;default:null"`
	Quantity  int       `json:"quantity" gorm:"column:quantity;default:null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
