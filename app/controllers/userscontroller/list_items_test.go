package userscontroller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestUsersController_ListItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockItemService := mock.NewMockItemServiceInterface(ctrl)

	usersController := NewUsersController()
	usersController.itemService = mockItemService

	router := httprouter.New()
	router.GET("/items", usersController.ListItems)

	t.Run("ListItems_Success", func(t *testing.T) {
		mockItems := []*models.Item{
			{ID: uuid.NewString(), Name: "Item 1", Price: 10.0},
			{ID: uuid.NewString(), Name: "Item 2", Price: 15.0},
		}

		mockItemService.EXPECT().ListItems().Return(mockItems, nil)

		request := httptest.NewRequest("GET", "/items", nil)
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)

		var response controllers.Response
		err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.Empty(t, response.Error)

		actualItems := []*models.Item{}
		respBytes, _ := json.Marshal(response.Data)
		err = json.Unmarshal(respBytes, &actualItems)
		assert.Empty(t, response.Error)
		assert.Equal(t, mockItems, actualItems)
	})

	t.Run("ListItems_Error", func(t *testing.T) {
		mockItemService.EXPECT().ListItems().Return(nil, errors.New("failed to fetch items"))

		request := httptest.NewRequest("GET", "/items", nil)
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)

		var response controllers.Response
		err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.False(t, response.Success)
		assert.Equal(t, response.Error, "failed to fetch items")
	})
}
