package repository

import (
	"fmt"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/config/db"
	"github.com/jinzhu/gorm"
)

type UserRepositoryInterface interface {
	Create(user *models.User) error
	Find(userId string) (*models.User, error)
	FindByName(username string) (*models.User, error)
	Suspend(user *models.User) error
}

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository returns instance of UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{db: db.Get()}
}

func (ur *UserRepository) Create(user *models.User) error {
	if err := ur.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Find(userId string) (*models.User, error) {
	user := &models.User{}

	err := ur.db.Where("users.id = ?", userId).Find(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("No user exists for userId : %v", userId)
	}

	return user, nil
}

func (ur *UserRepository) FindByName(username string) (*models.User, error) {
	user := &models.User{}

	err := ur.db.Where("users.username = ?", username).Find(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("No user exists for username : %v", username)
	}

	return user, err
}

func (ur *UserRepository) Suspend(user *models.User) error {
	user.Active = false
	if err := ur.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}
