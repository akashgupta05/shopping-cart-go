package userscontroller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/akashgupta05/shopping-cart-go/app/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestUsersController_AddToCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock.NewMockUserServiceInterface(ctrl)

	usersController := NewUsersController()
	usersController.userService = mockUserService

	router := httprouter.New()
	router.POST("/users/cart/add", usersController.AddToCart)

	t.Run("AddToCart_Success", func(t *testing.T) {
		mockUserID := "123"
		mockItemID := "item1"
		mockQuantity := 2

		mockUserService.EXPECT().AddToCart(mockUserID, mockItemID, mockQuantity).Return(nil)

		payload := `{"item_id":"item1","quantity":2}`
		request := httptest.NewRequest("POST", "/users/cart/add?user_id="+mockUserID, strings.NewReader(payload))
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)

		var response controllers.Response
		err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.Empty(t, response.Error)
	})

	t.Run("AddToCart_InvalidRequest", func(t *testing.T) {
		mockUserID := "123"
		payload := `{"item_id":"","quantity":0}` // Invalid payload with missing item_id and quantity

		request := httptest.NewRequest("POST", "/users/cart/add?user_id="+mockUserID, strings.NewReader(payload))
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

		var response controllers.Response
		err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.False(t, response.Success)
		assert.Contains(t, response.Error, "missing item_id")
	})

	t.Run("AddToCart_Error", func(t *testing.T) {
		mockUserID := "123"
		mockItemID := "item1"
		mockQuantity := 2

		mockUserService.EXPECT().AddToCart(mockUserID, mockItemID, mockQuantity).Return(errors.New("failed to add item to cart"))

		payload := `{"item_id":"item1","quantity":2}`
		request := httptest.NewRequest("POST", "/users/cart/add?user_id="+mockUserID, strings.NewReader(payload))
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)

		var response controllers.Response
		err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.False(t, response.Success)
		assert.Contains(t, response.Error, "failed to add item to cart")
	})
}
