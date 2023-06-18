package userscontroller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/akashgupta05/shopping-cart-go/app/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestUsersController_RemoveItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock.NewMockUserServiceInterface(ctrl)

	usersController := NewUsersController()
	usersController.userService = mockUserService

	router := httprouter.New()
	router.DELETE("/users/cart/:item_id", usersController.RemoveItem)

	t.Run("RemoveItem_Success", func(t *testing.T) {
		mockUserID := uuid.NewString()
		mockItemID := uuid.NewString()

		mockUserService.EXPECT().RemoveFromCart(mockUserID, mockItemID).Return(nil)

		request := httptest.NewRequest("DELETE", "/users/cart/"+mockItemID+"?user_id="+mockUserID, nil)
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)

		var response controllers.Response
		err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.Empty(t, response.Error)
	})

	t.Run("RemoveItem_Error", func(t *testing.T) {
		mockUserID := uuid.NewString()
		mockItemID := uuid.NewString()

		mockUserService.EXPECT().RemoveFromCart(mockUserID, mockItemID).Return(errors.New("failed to remove item from cart"))

		request := httptest.NewRequest("DELETE", "/users/cart/"+mockItemID+"?user_id="+mockUserID, nil)
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)

		var response controllers.Response
		err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.False(t, response.Success)
		assert.Contains(t, response.Error, "failed to remove item from cart")
	})
}
