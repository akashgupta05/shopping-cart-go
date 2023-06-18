package services

import (
	"fmt"
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/repository/mock"
	"github.com/akashgupta05/shopping-cart-go/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type AccessTokenMatcher struct {
	UserID string
	Active bool
}

func (m *AccessTokenMatcher) Matches(x interface{}) bool {
	token, ok := x.(*models.AccessToken)
	if !ok {
		return false
	}

	return token.UserID == m.UserID && token.Active == m.Active && token.Token != ""
}

func (m *AccessTokenMatcher) String() string {
	return fmt.Sprintf("matches access token: %s", m.UserID)
}

func TestAuthService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepositoryInterface(ctrl)
	mockAccessTokenRepo := mock.NewMockAccessTokensRepositoryInterface(ctrl)
	mockCartSessionRepo := mock.NewMockCartSessionsRepositoryInterface(ctrl)

	authService := &AuthService{
		userRepo:        mockUserRepo,
		accessTokenRepo: mockAccessTokenRepo,
		cartSessionRepo: mockCartSessionRepo,
	}

	passDigest, err := utils.GeneratePasswordDigest("password")
	assert.Nil(t, err)

	mockUser := &models.User{
		ID:             uuid.NewString(),
		Username:       "testuser",
		PasswordDigest: passDigest,
		Role:           models.USER,
		Active:         true,
	}
	mockUserRepo.EXPECT().FindByName("testuser").Return(mockUser, nil)

	mockAccessToken := &AccessTokenMatcher{
		UserID: mockUser.ID,
		Active: true,
	}
	mockAccessTokenRepo.EXPECT().Upsert(mockAccessToken).Return(nil)

	mockCartSessionRepo.EXPECT().GetByUserID(mockUser.ID).Return(nil, nil)
	mockCartSession := &models.CartSession{
		UserID: mockUser.ID,
		Status: models.ACTIVE,
	}
	mockCartSessionRepo.EXPECT().Upsert(mockCartSession).Return(nil)

	token, err := authService.Login("testuser", "password")

	assert.NoError(t, err, "Login should not return an error")
	assert.NotEmpty(t, token, "Access token should not be empty")
}

func TestAuthService_Logout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccessTokenRepo := mock.NewMockAccessTokensRepositoryInterface(ctrl)

	authService := &AuthService{
		accessTokenRepo: mockAccessTokenRepo,
	}

	mockAccessToken := "token123"
	mockAccessTokenRepo.EXPECT().MarkInactive(mockAccessToken).Return(nil)

	err := authService.Logout(mockAccessToken)

	assert.NoError(t, err, "Logout should not return an error")
}

func TestAuthService_ValidateAccessToken(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccessTokenRepo := mock.NewMockAccessTokensRepositoryInterface(ctrl)
	mockUserRepo := mock.NewMockUserRepositoryInterface(ctrl)

	authService := &AuthService{
		accessTokenRepo: mockAccessTokenRepo,
		userRepo:        mockUserRepo,
	}

	mockAccessToken := "token123"
	mockUserID := "user123"
	mockAccessTokenRepo.EXPECT().ValidateToken(mockAccessToken).Return(mockUserID, nil)

	mockUser := &models.User{
		ID:   mockUserID,
		Role: models.USER,
	}
	mockUserRepo.EXPECT().Find(mockUserID).Return(mockUser, nil)

	valid, userID := authService.ValidateAccessToken(mockAccessToken, "user")

	assert.True(t, valid, "Access token should be valid")
	assert.Equal(t, mockUserID, userID, "User ID should match")
}
