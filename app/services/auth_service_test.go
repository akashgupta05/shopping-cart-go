package services

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/repository/mock"
	"github.com/akashgupta05/shopping-cart-go/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_LoginWithJWT(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepositoryInterface(ctrl)
	mockCartSessionRepo := mock.NewMockCartSessionsRepositoryInterface(ctrl)

	authService := AuthService{
		userRepo:        mockUserRepo,
		cartSessionRepo: mockCartSessionRepo,
	}

	username := "testuser"
	password := "testpassword"
	expectedUserID := "12345"
	expectedRole := string(models.USER)
	passDigest, err := utils.GeneratePasswordDigest("testpassword")
	assert.Nil(t, err)

	t.Run("Valid credentials", func(t *testing.T) {
		user := &models.User{
			ID:             expectedUserID,
			Role:           models.Role(expectedRole),
			Active:         true,
			PasswordDigest: passDigest,
		}

		mockUserRepo.EXPECT().FindByName(username).Return(user, nil)
		mockCartSessionRepo.EXPECT().GetByUserID(expectedUserID).Return(&models.CartSession{}, nil)

		tokenString, expirationTime, err := authService.LoginWithJWT(username, password)

		assert.NoError(t, err)
		assert.NotNil(t, tokenString)
		assert.NotNil(t, expirationTime)
	})

	t.Run("Valid credentials and cart session doesn't exist", func(t *testing.T) {
		user := &models.User{
			ID:             expectedUserID,
			Role:           models.Role(expectedRole),
			Active:         true,
			PasswordDigest: passDigest,
		}

		mockUserRepo.EXPECT().FindByName(username).Return(user, nil)
		mockCartSessionRepo.EXPECT().GetByUserID(expectedUserID).Return(nil, nil)
		mockCartSessionRepo.EXPECT().Upsert(gomock.Any()).Return(nil)

		tokenString, expirationTime, err := authService.LoginWithJWT(username, password)

		assert.NoError(t, err)
		assert.NotNil(t, tokenString)
		assert.NotNil(t, *expirationTime)
	})

	t.Run("Invalid username and password", func(t *testing.T) {
		mockUserRepo.EXPECT().FindByName(username).Return(nil, errors.New("user not found"))

		tokenString, expirationTime, err := authService.LoginWithJWT(username, password)

		assert.EqualError(t, err, fmt.Sprintf("Failed to validate username and password: Failed to fetch user: user not found"))
		assert.Empty(t, tokenString)
		assert.Nil(t, expirationTime)
	})

	t.Run("Suspended user", func(t *testing.T) {
		user := &models.User{
			ID:             expectedUserID,
			Role:           models.Role(expectedRole),
			Active:         false,
			PasswordDigest: passDigest,
		}

		mockUserRepo.EXPECT().FindByName(username).Return(user, nil)

		tokenString, expirationTime, err := authService.LoginWithJWT(username, password)

		assert.EqualError(t, err, "Failed to validate username and password: Suspended User")
		assert.Empty(t, tokenString)
		assert.Nil(t, expirationTime)
	})

	t.Run("Invalid password", func(t *testing.T) {
		user := &models.User{
			ID:             expectedUserID,
			Role:           models.Role(expectedRole),
			Active:         true,
			PasswordDigest: passDigest,
		}

		mockUserRepo.EXPECT().FindByName(username).Return(user, nil)

		tokenString, expirationTime, err := authService.LoginWithJWT(username, "wrong_password")

		assert.EqualError(t, err, "Failed to validate username and password: Invalid Password")
		assert.Empty(t, tokenString)
		assert.Nil(t, expirationTime)
	})
}

func TestAuthService_ValidateJWT(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := AuthService{}
	role := string(models.USER)
	userID := uuid.NewString()
	expirationTime := time.Now().Add(JWT_TOKEN_EXPIRY)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(JWT_SECRET_KEY)
	assert.Nil(t, err)

	t.Run("Valid JWT token", func(t *testing.T) {
		valid, actualUserID := authService.ValidateJWT(jwtToken, role)

		assert.True(t, valid)
		assert.Equal(t, userID, actualUserID)
	})

	t.Run("Invalid JWT token", func(t *testing.T) {
		valid, actualUserID := authService.ValidateJWT("invalid_token", role)

		assert.False(t, valid)
		assert.Empty(t, actualUserID)
	})
}
