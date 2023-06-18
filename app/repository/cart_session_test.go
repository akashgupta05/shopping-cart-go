package repository

import (
	"testing"
	"time"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/testhelpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCartSessionsRepository_Upsert(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := CartSessionsRepository{db: testhelpers.Get()}

	cartSession := &models.CartSession{
		UserID: uuid.NewString(),
		Total:  100,
		Status: models.ACTIVE,
	}

	err := repo.Upsert(cartSession)
	assert.NoError(t, err)
}

func TestCartSessionsRepository_GetByUserID(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := CartSessionsRepository{db: testhelpers.Get()}

	userID := uuid.NewString()

	activeCartSession := &models.CartSession{
		UserID:    userID,
		Total:     100,
		Status:    models.ACTIVE,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.Upsert(activeCartSession)
	assert.NoError(t, err)

	result, err := repo.GetByUserID(userID)
	assert.NoError(t, err)
	assert.Equal(t, activeCartSession.UserID, result.UserID)
	assert.NotEmpty(t, activeCartSession.Id)

	nonExistentUserID := uuid.NewString()
	result, err = repo.GetByUserID(nonExistentUserID)
	assert.NoError(t, err)
	assert.Nil(t, result)
}
