package repository

import (
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/testhelpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccessTokensRepository_Upsert(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := AccessTokensRepository{db: testhelpers.Get()}

	accessToken := &models.AccessToken{
		UserID: uuid.NewString(),
		Token:  "sample_token",
		Active: true,
	}

	// Test the Upsert method
	err := repo.Upsert(accessToken)
	assert.NoError(t, err)
}

func TestAccessTokensRepository_MarkInactive(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := AccessTokensRepository{db: testhelpers.Get()}

	accessToken := &models.AccessToken{
		UserID: uuid.NewString(),
		Token:  "sample_token1",
		Active: true,
	}

	err := repo.Upsert(accessToken)
	assert.NoError(t, err)

	err = repo.MarkInactive(accessToken.Token)
	assert.NoError(t, err)
}

func TestAccessTokensRepository_MarkInactiveForUser(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := AccessTokensRepository{db: testhelpers.Get()}

	accessToken := &models.AccessToken{
		UserID: uuid.NewString(),
		Token:  "sample_token",
		Active: true,
	}

	err := repo.Upsert(accessToken)
	assert.NoError(t, err)

	err = repo.MarkInactiveForUser(accessToken.UserID)
	assert.NoError(t, err)

	err = repo.MarkInactiveForUser("invalid user")
	assert.Error(t, err)
}

func TestAccessTokensRepository_ValidateToken(t *testing.T) {
	testhelpers.ConnectToTestDB(t)
	defer testhelpers.Close()

	repo := AccessTokensRepository{db: testhelpers.Get()}

	accessToken := &models.AccessToken{
		UserID: uuid.NewString(),
		Token:  "sample_token",
		Active: true,
	}

	err := repo.Upsert(accessToken)
	assert.NoError(t, err)

	userID, err := repo.ValidateToken(accessToken.Token)
	assert.NoError(t, err)
	assert.Equal(t, accessToken.UserID, userID)

	userID, err = repo.ValidateToken("Invalid token")
	assert.Error(t, err)
	assert.Equal(t, "", userID)
}
