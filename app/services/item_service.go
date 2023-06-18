package services

import (
	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/repository"
)

type ItemServiceInterface interface {
	ListItems() ([]*models.Item, error)
}

type ItemService struct {
	itemRepo repository.ItemRepositoryInterface
}

func NewItemService() *ItemService {
	return &ItemService{
		itemRepo: repository.NewItemRepository(),
	}
}

func (is *ItemService) ListItems() ([]*models.Item, error) {
	items, err := is.itemRepo.List()
	if err != nil {
		return nil, err
	}

	return items, nil
}
