package repository

import (
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/testhelpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCartItemsRepository_Upsert(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := CartItemsRepository{db: testhelpers.Get()}

	cartItem := &models.CartItem{
		ItemID:    uuid.NewString(),
		SessionID: uuid.NewString(),
		Quantity:  1,
	}

	err := repo.Upsert(cartItem)
	assert.NoError(t, err)
}

func TestCartItemsRepository_GetByItemID(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := CartItemsRepository{db: testhelpers.Get()}

	itemID := uuid.NewString()
	sessionID := uuid.NewString()

	cartItem := &models.CartItem{
		ItemID:    itemID,
		SessionID: sessionID,
		Quantity:  1,
	}

	err := repo.Upsert(cartItem)
	assert.NoError(t, err)

	result, err := repo.GetByItemID(sessionID, itemID)
	assert.NoError(t, err)
	assert.Equal(t, cartItem.ItemID, result.ItemID)
	assert.Equal(t, cartItem.SessionID, result.SessionID)

	nonExistentItemID := uuid.NewString()
	result, err = repo.GetByItemID(sessionID, nonExistentItemID)
	assert.NoError(t, err)
}

func TestCartItemsRepository_List(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := CartItemsRepository{db: testhelpers.Get()}

	sessionID := uuid.NewString()

	cartItems := []*models.CartItem{
		{
			ItemID:    uuid.NewString(),
			SessionID: sessionID,
			Quantity:  1,
		},
		{
			ItemID:    uuid.NewString(),
			SessionID: sessionID,
			Quantity:  2,
		},
	}

	for _, item := range cartItems {
		err := repo.Upsert(item)
		assert.NoError(t, err)
	}

	result, err := repo.List(sessionID)
	assert.NoError(t, err)
	assert.Equal(t, len(cartItems), len(result))

	nonExistentSessionID := uuid.NewString()
	result, err = repo.List(nonExistentSessionID)
	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestCartItemsRepository_Delete(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := CartItemsRepository{db: testhelpers.Get()}

	itemID := uuid.NewString()
	sessionID := uuid.NewString()

	cartItem := &models.CartItem{
		ItemID:    itemID,
		SessionID: sessionID,
		Quantity:  1,
	}

	err := repo.Upsert(cartItem)
	assert.NoError(t, err)

	err = repo.Delete(sessionID, itemID)
	assert.NoError(t, err)

	result, err := repo.GetByItemID(sessionID, itemID)
	assert.NoError(t, err)
	assert.Empty(t, result)

	nonExistentItemID := uuid.NewString()
	err = repo.Delete(sessionID, nonExistentItemID)
	assert.NoError(t, err)
}
