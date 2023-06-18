package services

//go:generate mockgen -source=admin_service.go -destination=mock/mock_admin_service.go -package=mock

import (
	"errors"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/repository"
	"github.com/sirupsen/logrus"
)

type AdminServiceInterface interface {
	SuspendUser(userID string) error
	AddItems(items []*models.Item) ([]*models.Item, error)
}

type AdminService struct {
	userRepo         repository.UserRepositoryInterface
	itemRepo         repository.ItemRepositoryInterface
	accessTokensRepo repository.AccessTokensRepositoryInterface
}

// NewAdminService returns instance of AdminService
func NewAdminService() *AdminService {
	return &AdminService{
		userRepo:         repository.NewUserRepository(),
		itemRepo:         repository.NewItemRepository(),
		accessTokensRepo: repository.NewAccessTokensRepository(),
	}
}

func (s *AdminService) SuspendUser(userID string) error {
	user, err := s.userRepo.Find(userID)
	if err != nil {
		return err
	}

	if user.Role == models.ADMIN {
		return errors.New("Cannot suspend admin user")
	}

	err = s.accessTokensRepo.MarkInactiveForUser(userID)
	if err != nil {
		logrus.Error("Failed to mark access tokens inactive for user", userID, err.Error())
	}

	err = s.userRepo.Suspend(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AdminService) AddItems(items []*models.Item) ([]*models.Item, error) {
	addedItems := []*models.Item{}

	for _, item := range items {
		currentItem, err := s.itemRepo.GetByName(item.Name)
		if err != nil {
			logrus.Warn("Failed to get item info", item.Name)
			continue
		}

		if currentItem != nil {
			currentItem.Quantity += item.Quantity
			err = s.itemRepo.Upsert(currentItem)
			if err != nil {
				logrus.Warn("Failed to update item", item.Name)
				continue
			}

			addedItems = append(addedItems, currentItem)
			continue
		}

		err = s.itemRepo.Upsert(item)
		if err != nil {
			logrus.Warn("Failed to save item", item.Name)
			continue
		}
		addedItems = append(addedItems, item)
	}

	return addedItems, nil
}
