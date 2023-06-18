package authcontroller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"

	"github.com/akashgupta05/shopping-cart-go/app/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestAuthController_LogoutUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mock.NewMockAuthServiceInterface(ctrl)

	authController := NewAuthController()
	authController.authService = mockAuthService

	router := httprouter.New()
	router.POST("/logout", authController.LogoutUser)

	t.Run("LogoutUser_Success", func(t *testing.T) {
		accessToken := "mockAccessToken"

		mockAuthService.EXPECT().Logout(accessToken).Return(nil)

		request := httptest.NewRequest("POST", "/logout", nil)
		request.Header.Set("Access-Token", accessToken)
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)

		var response controllers.Response
		_ = json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.True(t, response.Success)
		assert.Empty(t, response.Error)
	})

	t.Run("LogoutUser_InvalidToken", func(t *testing.T) {
		accessToken := "mockInvalidToken"

		mockAuthService.EXPECT().Logout(accessToken).Return(errors.New("invalid access token"))

		request := httptest.NewRequest("POST", "/logout", nil)
		request.Header.Set("Access-Token", accessToken)
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

		var response controllers.Response
		_ = json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.False(t, response.Success)
		assert.Contains(t, response.Error, "invalid access token")
	})
}
