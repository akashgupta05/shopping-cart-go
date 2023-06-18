package repository

import (
	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/config/db"
	"github.com/jinzhu/gorm"
)

type AccessTokensRepositoryInterface interface {
	Upsert(accessToken *models.AccessToken) error
	MarkInactive(token string) error
	MarkInactiveForUser(userID string) error
	ValidateToken(token string) (string, error)
}

type AccessTokensRepository struct {
	db *gorm.DB
}

// NewAccessTokensRepository returns instance of AccessTokensRepository
func NewAccessTokensRepository() *AccessTokensRepository {
	return &AccessTokensRepository{db: db.Get()}
}

func (atr *AccessTokensRepository) Upsert(accessToken *models.AccessToken) error {
	if err := atr.db.Save(accessToken).Error; err != nil {
		return err
	}

	return nil
}

func (atr *AccessTokensRepository) MarkInactive(token string) error {
	if err := atr.db.Model(&models.AccessToken{}).Where("token = ?", token).Update("active", false).Error; err != nil {
		return err
	}

	return nil
}

func (atr *AccessTokensRepository) MarkInactiveForUser(userID string) error {
	if err := atr.db.Model(&models.AccessToken{}).Where("user_id = ?", userID).Update("active", false).Error; err != nil {
		return err
	}

	return nil
}

func (atr *AccessTokensRepository) ValidateToken(token string) (string, error) {
	accessToken := &models.AccessToken{}
	if err := atr.db.Where("access_tokens.token = ? and access_tokens.active = true", token).Find(accessToken).Error; err != nil {
		return "", err
	}

	return accessToken.UserID, nil
}
