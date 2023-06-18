package models

import "time"

type Item struct {
	Id        string    `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string    `json:"name" gorm:"not null;column:name;default:null"`
	Quantity  int       `json:"quantity" gorm:"column:quantity;default:null"`
	Price     int       `json:"price" gorm:"column:price;default:null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
