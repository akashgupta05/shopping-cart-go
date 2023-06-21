package services

//go:generate mockgen -source=auth_service.go -destination=mock/mock_auth_service.go -package=mock

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/repository"
	"github.com/akashgupta05/shopping-cart-go/utils"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))
var JWT_TOKEN_EXPIRY = 10 * time.Minute

type AuthServiceInterface interface {
	LoginWithJWT(username, password string) (string, *time.Time, error)
	ValidateJWT(jwtToken, role string) (bool, string)
	RefreshJWT(userID string) (string, *time.Time, error)
}

type AuthService struct {
	userRepo        repository.UserRepositoryInterface
	cartSessionRepo repository.CartSessionsRepositoryInterface
}

// NewAuthService returns instance of AuthService
func NewAuthService() *AuthService {
	return &AuthService{
		userRepo:        repository.NewUserRepository(),
		cartSessionRepo: repository.NewCartSessionsRepository(),
	}
}

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (s *AuthService) ValidateJWT(jwtToken, role string) (bool, string) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET_KEY, nil
	})

	if err != nil || !token.Valid {
		return false, ""
	}

	return models.Role(claims.Role) == models.Role(role), claims.UserID
}

func (s *AuthService) LoginWithJWT(username, password string) (string, *time.Time, error) {
	user, err := s.validateUsernamePassword(username, password)
	if err != nil {
		return "", nil, fmt.Errorf("Failed to validate username and password: %s", err.Error())
	}

	token, expiresAt := s.generateJWT(user)
	tokenString, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return "", nil, fmt.Errorf("Failed to generate token %s", err.Error())
	}

	cartSession, err := s.cartSessionRepo.GetByUserID(user.ID)
	if err != nil {
		return "", nil, nil
	}

	if cartSession != nil {
		return tokenString, expiresAt, nil
	}

	cartSession = &models.CartSession{
		UserID: user.ID,
		Status: models.ACTIVE,
	}

	if err := s.cartSessionRepo.Upsert(cartSession); err != nil {
		return tokenString, expiresAt, fmt.Errorf("Failed to create cart session %s", err.Error())
	}

	return tokenString, expiresAt, nil
}

func (s *AuthService) RefreshJWT(userID string) (string, *time.Time, error) {
	user, err := s.userRepo.Find(userID)
	if err != nil {
		return "", nil, fmt.Errorf("Failed to validate username and password: %s", err.Error())
	}

	token, expiresAt := s.generateJWT(user)
	tokenString, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return "", nil, fmt.Errorf("Failed to generate token %s", err.Error())
	}

	return tokenString, expiresAt, nil
}

func (s *AuthService) validateUsernamePassword(username, password string) (*models.User, error) {
	user, err := s.userRepo.FindByName(username)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch user: %s", err.Error())
	}

	if !user.Active {
		return nil, errors.New("Suspended User")
	}

	if !utils.ComparePasswordDigest(password, user.PasswordDigest) {
		return nil, fmt.Errorf("Invalid Password")
	}

	return user, nil
}

func (s *AuthService) generateJWT(user *models.User) (*jwt.Token, *time.Time) {
	expirationTime := time.Now().Add(JWT_TOKEN_EXPIRY)
	claims := &Claims{
		UserID: user.ID,
		Role:   string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims), &expirationTime
}
