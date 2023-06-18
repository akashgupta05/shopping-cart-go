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

func TestUsersController_GetCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock.NewMockUserServiceInterface(ctrl)

	usersController := NewUsersController()
	usersController.userService = mockUserService

	router := httprouter.New()
	router.GET("/users/cart", usersController.GetCart)

	t.Run("GetCart_Success", func(t *testing.T) {
		mockUserID := "123"
		mockCartItems := []*models.CartItem{
			{ID: uuid.NewString(), ItemID: uuid.NewString(), SessionID: uuid.NewString(), Quantity: 2},
			{ID: uuid.NewString(), ItemID: uuid.NewString(), SessionID: uuid.NewString(), Quantity: 1},
		}

		mockUserService.EXPECT().GetCartItems(mockUserID).Return(mockCartItems, nil)

		request := httptest.NewRequest("GET", "/users/cart?user_id="+mockUserID, nil)
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)

		var response controllers.Response
		err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.Empty(t, response.Error)

		actualItems := []*models.CartItem{}
		respBytes, _ := json.Marshal(response.Data)
		err = json.Unmarshal(respBytes, &actualItems)
		assert.Empty(t, response.Error)
		assert.Equal(t, mockCartItems, actualItems)
	})

	t.Run("GetCart_Error", func(t *testing.T) {
		mockUserID := "123"
		mockUserService.EXPECT().GetCartItems(mockUserID).Return(nil, errors.New("failed to get cart items"))

		request := httptest.NewRequest("GET", "/users/cart?user_id="+mockUserID, nil)
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)

		var response controllers.Response
		err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.False(t, response.Success)
		assert.Contains(t, response.Error, "failed to get cart items")
	})
}
