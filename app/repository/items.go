package repository

import (
	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/config/db"
	"github.com/jinzhu/gorm"
)

type ItemRepositoryInterface interface {
	Upsert(item *models.Item) error
	Find(id string) (*models.Item, error)
	GetByName(itemName string) (*models.Item, error)
	List() ([]*models.Item, error)
}

type ItemRepository struct {
	db *gorm.DB
}

// NewItemRepository returns instance of ItemRepository
func NewItemRepository() *ItemRepository {
	return &ItemRepository{db: db.Get()}
}

func (ir *ItemRepository) Upsert(item *models.Item) error {
	if err := ir.db.Save(item).Error; err != nil {
		return err
	}

	return nil
}

func (ir *ItemRepository) Find(id string) (*models.Item, error) {
	item := &models.Item{}
	if err := ir.db.Where("items.id = ?", id).Find(item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (ir *ItemRepository) GetByName(itemName string) (*models.Item, error) {
	item := models.Item{}
	if err := ir.db.Where("items.name = ?", itemName).Find(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}

func (ir *ItemRepository) List() ([]*models.Item, error) {
	items := []*models.Item{}

	if err := ir.db.Where("items.quantity > 0").Find(&items).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return items, nil
}
