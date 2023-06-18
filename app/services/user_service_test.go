package services

import (
	"errors"
	"fmt"
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/repository/mock"
	"github.com/akashgupta05/shopping-cart-go/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type UserMatcher struct {
	Username       string
	Role           models.Role
	PasswordDigest string
	Active         bool
}

func (m *UserMatcher) Matches(x interface{}) bool {
	user, ok := x.(*models.User)
	if !ok {
		return false
	}

	return user.Username == m.Username &&
		user.Role == m.Role && user.Active == m.Active
}

func (m *UserMatcher) String() string {
	return fmt.Sprintf("matches access user: %s", m.Username)
}

func TestUserService_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepositoryInterface(ctrl)
	mockItemRepo := mock.NewMockItemRepositoryInterface(ctrl)
	mockCartItemRepo := mock.NewMockCartItemsRepositoryInterface(ctrl)
	mockCartSessionRepo := mock.NewMockCartSessionsRepositoryInterface(ctrl)

	userService := UserService{
		userRepo:        mockUserRepo,
		itemRepo:        mockItemRepo,
		cartItemRepo:    mockCartItemRepo,
		cartSessionRepo: mockCartSessionRepo,
	}

	userID := "123"
	mockUser := &models.User{
		ID:       uuid.NewString(),
		Username: "john",
		Role:     models.USER,
		Active:   true,
	}

	mockUserRepo.EXPECT().Find(userID).Return(mockUser, nil)

	user, err := userService.GetUserByID(userID)
	assert.NoError(t, err, "GetUserByID should not return an error")
	assert.Equal(t, mockUser, user, "GetUserByID should return the mock user")
}

func TestUserService_GetUserByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepositoryInterface(ctrl)
	mockItemRepo := mock.NewMockItemRepositoryInterface(ctrl)
	mockCartItemRepo := mock.NewMockCartItemsRepositoryInterface(ctrl)
	mockCartSessionRepo := mock.NewMockCartSessionsRepositoryInterface(ctrl)

	userService := UserService{
		userRepo:        mockUserRepo,
		itemRepo:        mockItemRepo,
		cartItemRepo:    mockCartItemRepo,
		cartSessionRepo: mockCartSessionRepo,
	}

	userID := "123"
	mockUserRepo.EXPECT().Find(userID).Return(nil, errors.New("user not found"))

	user, err := userService.GetUserByID(userID)
	assert.Error(t, err, "GetUserByID should return an error")
	assert.Nil(t, user, "GetUserByID should return nil user")
}

func TestUserService_AddToCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepositoryInterface(ctrl)
	mockItemRepo := mock.NewMockItemRepositoryInterface(ctrl)
	mockCartItemRepo := mock.NewMockCartItemsRepositoryInterface(ctrl)
	mockCartSessionRepo := mock.NewMockCartSessionsRepositoryInterface(ctrl)

	userService := UserService{
		userRepo:        mockUserRepo,
		itemRepo:        mockItemRepo,
		cartItemRepo:    mockCartItemRepo,
		cartSessionRepo: mockCartSessionRepo,
	}

	userID := "123"
	itemID := "456"
	quantity := 2

	mockItem := &models.Item{
		ID:       uuid.NewString(),
		Name:     "Item 1",
		Quantity: 5,
	}

	mockCartSession := &models.CartSession{
		ID:     uuid.NewString(),
		UserID: userID,
		Status: models.ACTIVE,
	}

	mockItemRepo.EXPECT().Find(itemID).Return(mockItem, nil)
	mockCartSessionRepo.EXPECT().GetByUserID(userID).Return(mockCartSession, nil)
	mockCartItemRepo.EXPECT().GetByItemID(mockCartSession.ID, itemID).Return(nil, nil)
	mockCartItemRepo.EXPECT().Upsert(gomock.Any()).Return(nil)

	err := userService.AddToCart(userID, itemID, quantity)

	assert.NoError(t, err, "AddToCart should not return an error")
}

func TestUserService_AddToCart_InsufficientStock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepositoryInterface(ctrl)
	mockItemRepo := mock.NewMockItemRepositoryInterface(ctrl)
	mockCartItemRepo := mock.NewMockCartItemsRepositoryInterface(ctrl)
	mockCartSessionRepo := mock.NewMockCartSessionsRepositoryInterface(ctrl)

	userService := UserService{
		userRepo:        mockUserRepo,
		itemRepo:        mockItemRepo,
		cartItemRepo:    mockCartItemRepo,
		cartSessionRepo: mockCartSessionRepo,
	}

	userID := "123"
	itemID := "456"
	quantity := 10

	mockItem := &models.Item{
		ID:       uuid.NewString(),
		Name:     "Item 1",
		Quantity: 5,
	}

	// mockCartSession := &models.CartSession{
	// 	ID:     uuid.NewString(),
	// 	UserID: userID,
	// 	Status: models.ACTIVE,
	// }

	mockItemRepo.EXPECT().Find(itemID).Return(mockItem, nil)

	err := userService.AddToCart(userID, itemID, quantity)

	assert.Error(t, err, "AddToCart should return an error")
	assert.EqualError(t, err, "Insufficient stock", "AddToCart should return an error message")
}

func TestUserService_RemoveFromCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepositoryInterface(ctrl)
	mockItemRepo := mock.NewMockItemRepositoryInterface(ctrl)
	mockCartItemRepo := mock.NewMockCartItemsRepositoryInterface(ctrl)
	mockCartSessionRepo := mock.NewMockCartSessionsRepositoryInterface(ctrl)

	userService := UserService{
		userRepo:        mockUserRepo,
		itemRepo:        mockItemRepo,
		cartItemRepo:    mockCartItemRepo,
		cartSessionRepo: mockCartSessionRepo,
	}

	userID := "123"
	itemID := "456"

	mockCartSession := &models.CartSession{
		ID:     uuid.NewString(),
		UserID: userID,
		Status: models.ACTIVE,
	}

	mockCartSessionRepo.EXPECT().GetByUserID(userID).Return(mockCartSession, nil)
	mockCartItemRepo.EXPECT().Delete(mockCartSession.ID, itemID).Return(nil)

	err := userService.RemoveFromCart(userID, itemID)

	assert.NoError(t, err, "RemoveFromCart should not return an error")
}

func TestUserService_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepositoryInterface(ctrl)
	mockItemRepo := mock.NewMockItemRepositoryInterface(ctrl)
	mockCartItemRepo := mock.NewMockCartItemsRepositoryInterface(ctrl)
	mockCartSessionRepo := mock.NewMockCartSessionsRepositoryInterface(ctrl)

	userService := UserService{
		userRepo:        mockUserRepo,
		itemRepo:        mockItemRepo,
		cartItemRepo:    mockCartItemRepo,
		cartSessionRepo: mockCartSessionRepo,
	}

	passDigest, err := utils.GeneratePasswordDigest("password")
	assert.Nil(t, err)

	mockUser := &UserMatcher{
		Username:       "john",
		Role:           models.USER,
		PasswordDigest: passDigest,
		Active:         true,
	}

	mockUserRepo.EXPECT().FindByName("john").Return(nil, nil)
	mockUserRepo.EXPECT().Create(mockUser).Return(nil)

	user, err := userService.RegisterUser("john", "password", models.USER)

	assert.NoError(t, err, "RegisterUser should not return an error")
	assert.Equal(t, mockUser.Username, user.Username, "RegisterUser should return the mock user")
	assert.Equal(t, mockUser.Active, user.Active)
}

func TestUserService_RegisterUser_DuplicateUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepositoryInterface(ctrl)
	mockItemRepo := mock.NewMockItemRepositoryInterface(ctrl)
	mockCartItemRepo := mock.NewMockCartItemsRepositoryInterface(ctrl)
	mockCartSessionRepo := mock.NewMockCartSessionsRepositoryInterface(ctrl)

	userService := UserService{
		userRepo:        mockUserRepo,
		itemRepo:        mockItemRepo,
		cartItemRepo:    mockCartItemRepo,
		cartSessionRepo: mockCartSessionRepo,
	}

	username := "john"
	password := "password"
	role := models.USER

	existingUser := &models.User{
		ID:       uuid.NewString(),
		Username: username,
		Role:     role,
	}

	mockUserRepo.EXPECT().FindByName(username).Return(existingUser, nil)

	user, err := userService.RegisterUser(username, password, role)

	assert.Error(t, err, "RegisterUser should return an error")
	assert.Nil(t, user, "RegisterUser should return nil user")
	assert.EqualError(t, err, "user already present with username: john", "RegisterUser should return an error message")
}

func TestUserService_GetCartItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepositoryInterface(ctrl)
	mockItemRepo := mock.NewMockItemRepositoryInterface(ctrl)
	mockCartItemRepo := mock.NewMockCartItemsRepositoryInterface(ctrl)
	mockCartSessionRepo := mock.NewMockCartSessionsRepositoryInterface(ctrl)

	userService := UserService{
		userRepo:        mockUserRepo,
		itemRepo:        mockItemRepo,
		cartItemRepo:    mockCartItemRepo,
		cartSessionRepo: mockCartSessionRepo,
	}

	userID := "123"

	mockCartSession := &models.CartSession{
		ID:     uuid.NewString(),
		UserID: userID,
		Status: models.ACTIVE,
	}

	mockCartItems := []*models.CartItem{
		{
			ID:       uuid.NewString(),
			ItemID:   "101",
			Quantity: 2,
		},
		{
			ID:       uuid.NewString(),
			ItemID:   "102",
			Quantity: 3,
		},
	}

	mockCartSessionRepo.EXPECT().GetByUserID(userID).Return(mockCartSession, nil)
	mockCartItemRepo.EXPECT().List(mockCartSession.ID).Return(mockCartItems, nil)

	cartItems, err := userService.GetCartItems(userID)
	assert.NoError(t, err, "GetCartItems should not return an error")
	assert.Equal(t, mockCartItems, cartItems, "GetCartItems should return the mock cart items")
}
