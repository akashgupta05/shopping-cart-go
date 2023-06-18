package services

//go:generate mockgen -source=user_service.go -destination=mock/mock_user_service.go -package=mock

import (
	"errors"
	"fmt"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/repository"
	"github.com/akashgupta05/shopping-cart-go/utils"
)

type UserServiceInterface interface {
	GetUserByID(userID string) (*models.User, error)
	AddToCart(userID, itemID string, quantity int) error
	RemoveFromCart(userID, itemID string) error
	RegisterUser(username, password string, role models.Role) (*models.User, error)
	GetCartItems(userID string) ([]*models.CartItem, error)
}

type UserService struct {
	userRepo        repository.UserRepositoryInterface
	itemRepo        repository.ItemRepositoryInterface
	cartItemRepo    repository.CartItemsRepositoryInterface
	cartSessionRepo repository.CartSessionsRepositoryInterface
}

func NewUserService() *UserService {
	return &UserService{
		userRepo:        repository.NewUserRepository(),
		itemRepo:        repository.NewItemRepository(),
		cartSessionRepo: repository.NewCartSessionsRepository(),
		cartItemRepo:    repository.NewCartItemsRepository(),
	}
}

func (s *UserService) GetUserByID(userID string) (*models.User, error) {
	user, err := s.userRepo.Find(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) AddToCart(userID, itemID string, quantity int) error {
	item, err := s.itemRepo.Find(itemID)
	if err != nil {
		return err
	}

	if item.Quantity < quantity {
		return errors.New("Insufficient stock")
	}

	cartSession, err := s.cartSessionRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	cartItem, err := s.cartItemRepo.GetByItemID(cartSession.ID, itemID)
	if err != nil {
		return err
	}

	if cartItem != nil {
		cartItem.Quantity = quantity
		err := s.cartItemRepo.Upsert(cartItem)
		if err != nil {
			return err
		}
		return nil
	}

	cartItem = &models.CartItem{
		ItemID:    itemID,
		Quantity:  quantity,
		SessionID: cartSession.ID,
	}

	err = s.cartItemRepo.Upsert(cartItem)
	if err != nil {
		return err
	}

	return nil
}

// RemoveFromCart removes an item from the user's cart
func (s *UserService) RemoveFromCart(userID, itemID string) error {
	cartSession, err := s.cartSessionRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	err = s.cartItemRepo.Delete(cartSession.ID, itemID)
	if err != nil {
		return err
	}

	return nil
}

// RemoveFromCart removes an item from the user's cart
func (s *UserService) RegisterUser(username, password string, role models.Role) (*models.User, error) {
	existingUser, _ := s.userRepo.FindByName(username)
	if existingUser != nil {
		return nil, fmt.Errorf("user already present with username: %s", username)
	}

	passwordDigest, err := utils.GeneratePasswordDigest(password)
	if err != nil {
		return nil, errors.New("Failed to register user")
	}

	user := &models.User{
		Username:       username,
		PasswordDigest: passwordDigest,
		Role:           role,
		Active:         true,
	}

	if err = s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("Failed to register user: %s", err.Error())
	}

	return user, nil
}

func (s *UserService) GetCartItems(userID string) ([]*models.CartItem, error) {
	cartSession, err := s.cartSessionRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	cartItems, err := s.cartItemRepo.List(cartSession.ID)
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}
