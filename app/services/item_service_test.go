package services

import (
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/repository/mock"
	"github.com/google/uuid"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestItemService_ListItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockItemRepo := mock.NewMockItemRepositoryInterface(ctrl)

	itemService := ItemService{
		itemRepo: mockItemRepo,
	}

	mockItems := []*models.Item{
		{
			ID:    uuid.NewString(),
			Name:  "Item 1",
			Price: 10,
		},
		{
			ID:    uuid.NewString(),
			Name:  "Item 2",
			Price: 19,
		},
	}

	mockItemRepo.EXPECT().List().Return(mockItems, nil)

	items, err := itemService.ListItems()

	assert.NoError(t, err, "ListItems should not return an error")
	assert.Equal(t, mockItems, items, "ListItems should return the mock items")
}
