package repository

//go:generate mockgen -source=cart_items.go -destination=./mock/mock_cart_items_repository.go -package=mock

import (
	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/config/db"
	"github.com/jinzhu/gorm"
)

type CartItemsRepositoryInterface interface {
	Upsert(cart *models.CartItem) error
	GetByItemID(sessionID, itemId string) (*models.CartItem, error)
	List(sessionID string) ([]*models.CartItem, error)
	Delete(sessionID, itemID string) error
}

type CartItemsRepository struct {
	db *gorm.DB
}

// NewCartItemsRepository returns instance of CartItemsRepository
func NewCartItemsRepository() *CartItemsRepository {
	return &CartItemsRepository{db: db.Get()}
}

func (cir *CartItemsRepository) Upsert(cartItem *models.CartItem) error {
	if err := cir.db.Save(cartItem).Error; err != nil {
		return err
	}

	return nil
}

func (cir *CartItemsRepository) Delete(sessionID, itemID string) error {
	if err := cir.db.Delete(&models.CartItem{}, "cart_items.session_id = ? and cart_items.item_id = ? ", sessionID, itemID).Error; err != nil {
		return err
	}

	return nil
}

func (cir *CartItemsRepository) GetByItemID(sessionID, itemId string) (*models.CartItem, error) {
	cartItem := &models.CartItem{}
	if err := cir.db.Where("cart_items.session_id = ? and cart_items.item_id = ?", sessionID, itemId).Find(cartItem).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return cartItem, nil
}

func (cir *CartItemsRepository) List(sessionID string) ([]*models.CartItem, error) {
	cartItems := []*models.CartItem{}

	if err := cir.db.Where("cart_items.session_id = ?", sessionID).Find(&cartItems).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return cartItems, nil
		}
		return nil, err
	}

	return cartItems, nil
}
