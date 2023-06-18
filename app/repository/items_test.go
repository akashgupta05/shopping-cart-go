package repository

import (
	"testing"
	"time"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/testhelpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestItemRepository_Upsert(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := ItemRepository{db: testhelpers.Get()}

	item := &models.Item{
		Name:     "Sample Item",
		Quantity: 10,
		Price:    100,
	}

	err := repo.Upsert(item)
	assert.NoError(t, err)
}

func TestItemRepository_Find(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := ItemRepository{db: testhelpers.Get()}

	item := &models.Item{
		Name:      "Sample Item",
		Quantity:  10,
		Price:     100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.Upsert(item)
	assert.NoError(t, err)

	result, err := repo.Find(item.ID)
	assert.NoError(t, err)
	assert.Equal(t, item.ID, result.ID)

	nonExistentItemID := uuid.NewString()
	result, err = repo.Find(nonExistentItemID)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestItemRepository_GetByName(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := ItemRepository{db: testhelpers.Get()}

	itemName := "Sample Item"

	item := &models.Item{
		Name:      itemName,
		Quantity:  10,
		Price:     100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.Upsert(item)
	assert.NoError(t, err)

	result, err := repo.GetByName(itemName)
	assert.NoError(t, err)
	assert.Equal(t, item.ID, result.ID)
	assert.Equal(t, item.Name, result.Name)

	nonExistentItemName := "Non-existent Item"
	result, err = repo.GetByName(nonExistentItemName)
	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestItemRepository_List(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := ItemRepository{db: testhelpers.Get()}

	items := []*models.Item{
		{
			Name:     "Item 1",
			Quantity: 10,
		},
		{
			Name:     "Item 2",
			Quantity: 5,
		},
		{
			Name:     "Item 3",
			Quantity: 4,
		},
	}

	for _, item := range items {
		err := repo.Upsert(item)
		assert.NoError(t, err)
	}

	result, err := repo.List()
	assert.NoError(t, err)
	assert.Equal(t, len(result), 3)
}
