package services

import (
	"errors"
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAdminService_SuspendUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepositoryInterface(ctrl)

	service := &AdminService{
		userRepo: mockUserRepo,
	}

	userID := "12345"
	user := &models.User{
		ID:   userID,
		Role: models.USER,
	}

	mockUserRepo.EXPECT().Find(userID).Return(user, nil)
	mockUserRepo.EXPECT().Suspend(user).Return(nil)

	err := service.SuspendUser(userID)
	assert.NoError(t, err)

	adminUserID := "admin123"
	adminUser := &models.User{
		ID:   adminUserID,
		Role: models.ADMIN,
	}

	mockUserRepo.EXPECT().Find(adminUserID).Return(adminUser, nil)

	err = service.SuspendUser(adminUserID)
	assert.Error(t, err)
	assert.Equal(t, "Cannot suspend admin user", err.Error())

	nonExistentUserID := "nonexistent123"

	mockUserRepo.EXPECT().Find(nonExistentUserID).Return(nil, errors.New("User not found"))

	err = service.SuspendUser(nonExistentUserID)
	assert.Error(t, err)
}

func TestAdminService_AddItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockItemRepo := mock.NewMockItemRepositoryInterface(ctrl)

	service := &AdminService{
		itemRepo: mockItemRepo,
	}

	items := []*models.Item{
		{
			ID:       uuid.NewString(),
			Name:     "Item 1",
			Quantity: 5,
		},
		{
			ID:       uuid.NewString(),
			Name:     "Item 2",
			Quantity: 3,
		},
	}

	mockItemRepo.EXPECT().GetByName("Item 1").Return(nil, nil)
	mockItemRepo.EXPECT().Upsert(items[0]).Return(nil)

	mockItemRepo.EXPECT().GetByName("Item 2").Return(nil, nil)
	mockItemRepo.EXPECT().Upsert(items[1]).Return(nil)

	addedItems, err := service.AddItems(items)
	assert.NoError(t, err)
	assert.Equal(t, items, addedItems)
}
