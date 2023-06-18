package repository

import (
	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/config/db"
	"github.com/jinzhu/gorm"
)

type CartSessionsRepositoryInterface interface {
	Upsert(cart *models.CartSession) error
	GetByUserID(userId string) (*models.CartSession, error)
}

type CartSessionsRepository struct {
	db *gorm.DB
}

// NewCartSessionsRepository returns instance of CartSessionsRepository
func NewCartSessionsRepository() *CartSessionsRepository {
	return &CartSessionsRepository{db: db.Get()}
}

func (csr *CartSessionsRepository) Upsert(cart *models.CartSession) error {
	if err := csr.db.Save(cart).Error; err != nil {
		return err
	}

	return nil
}

func (csr *CartSessionsRepository) GetByUserID(userId string) (*models.CartSession, error) {
	cartSession := &models.CartSession{}
	err := csr.db.Where("cart_sessions.user_id = ? and cart_sessions.status = ?", userId, models.ACTIVE).Find(cartSession).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return cartSession, nil
}
