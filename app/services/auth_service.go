package services

//go:generate mockgen -source=auth_service.go -destination=mock/mock_auth_service.go -package=mock

import (
	"errors"
	"fmt"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/repository"
	"github.com/akashgupta05/shopping-cart-go/utils"
)

type AuthServiceInterface interface {
	Login(username, password string) (string, error)
	Logout(accessToken string) error
	ValidateAccessToken(accessToken string, role string) (bool, string)
}

type AuthService struct {
	userRepo        repository.UserRepositoryInterface
	accessTokenRepo repository.AccessTokensRepositoryInterface
	cartSessionRepo repository.CartSessionsRepositoryInterface
}

// NewAuthService returns instance of AuthService
func NewAuthService() *AuthService {
	return &AuthService{
		userRepo:        repository.NewUserRepository(),
		accessTokenRepo: repository.NewAccessTokensRepository(),
		cartSessionRepo: repository.NewCartSessionsRepository(),
	}
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByName(username)
	if err != nil {
		return "", fmt.Errorf("Failed to fetch user: %s", err.Error())
	}

	if !user.Active {
		return "", errors.New("Suspended User")
	}

	if !utils.ComparePasswordDigest(password, user.PasswordDigest) {
		return "", fmt.Errorf("Invalid Password: %s", err.Error())
	}

	randomHex, err := utils.RandomHex()
	if err != nil {
		return "", fmt.Errorf("Failed to generate token: %s", err.Error())
	}

	accessToken := &models.AccessToken{
		UserID: user.ID,
		Active: true,
		Token:  randomHex,
	}
	if err := s.accessTokenRepo.Upsert(accessToken); err != nil {
		return "", nil
	}

	cartSession, err := s.cartSessionRepo.GetByUserID(user.ID)
	if err != nil {
		return "", nil
	}

	if cartSession != nil {
		return accessToken.Token, nil
	}

	cartSession = &models.CartSession{
		UserID: user.ID,
		Status: models.ACTIVE,
	}

	if err := s.cartSessionRepo.Upsert(cartSession); err != nil {
		return accessToken.Token, fmt.Errorf("Failed to create cart session %s", err.Error())
	}

	return accessToken.Token, nil
}

func (s *AuthService) Logout(accessToken string) error {
	if err := s.accessTokenRepo.MarkInactive(accessToken); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) ValidateAccessToken(accessToken string, role string) (bool, string) {
	userID, err := s.accessTokenRepo.ValidateToken(accessToken)
	if err != nil {
		return false, ""
	}

	user, err := s.userRepo.Find(userID)
	if err != nil {
		return false, ""
	}

	return user.Role == models.Role(role), user.ID
}
