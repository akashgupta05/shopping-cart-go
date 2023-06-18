package repository

import (
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/testhelpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := UserRepository{db: testhelpers.Get()}

	user := &models.User{
		Username:       "testuser",
		PasswordDigest: "password123",
		Role:           models.USER,
		Active:         true,
	}

	// Test the Create method
	err := repo.Create(user)
	assert.NoError(t, err)
}

func TestUserRepository_Find(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := UserRepository{db: testhelpers.Get()}

	user := &models.User{
		Username:       "testuser",
		PasswordDigest: "password123",
		Role:           models.USER,
		Active:         true,
	}

	err := repo.Create(user)
	assert.NoError(t, err)

	// Test the Find method
	result, err := repo.Find(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Username, result.Username)

	// Negative test case: Find a non-existent user
	nonExistentUserID := uuid.NewString()
	result, err = repo.Find(nonExistentUserID)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUserRepository_FindByName(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := UserRepository{db: testhelpers.Get()}

	username := "testuser"

	// Insert a sample user
	user := &models.User{
		Username:       username,
		PasswordDigest: "password123",
		Role:           models.USER,
		Active:         true,
	}

	err := repo.Create(user)
	assert.NoError(t, err)

	// Test the FindByName method
	result, err := repo.FindByName(username)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Username, result.Username)

	// Negative test case: Find a non-existent user by name
	nonExistentUsername := "nonexistentuser"
	result, err = repo.FindByName(nonExistentUsername)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUserRepository_Suspend(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := UserRepository{db: testhelpers.Get()}

	user := &models.User{
		Username:       "testuser",
		PasswordDigest: "password123",
		Role:           models.USER,
		Active:         true,
	}

	err := repo.Create(user)
	assert.NoError(t, err)

	// Test the Suspend method
	err = repo.Suspend(user)
	assert.NoError(t, err)
	assert.False(t, user.Active)
}
