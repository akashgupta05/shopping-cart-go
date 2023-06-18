package authcontroller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/akashgupta05/shopping-cart-go/app/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestAuthController_LoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mock.NewMockAuthServiceInterface(ctrl)

	authController := NewAuthController()
	authController.authService = mockAuthService

	router := httprouter.New()
	router.POST("/login", authController.LoginUser)

	t.Run("LoginUser_Success", func(t *testing.T) {
		loginPayload := &LoginPayload{
			Username: "testuser",
			Password: "testpassword",
		}

		mockAuthService.EXPECT().Login(loginPayload.Username, loginPayload.Password).Return("mockAccessToken", nil)

		payload, _ := json.Marshal(loginPayload)
		request := httptest.NewRequest("POST", "/login", bytes.NewBuffer(payload))
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)

		var response controllers.Response
		_ = json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.True(t, response.Success)
		assert.Empty(t, response.Error)
		assert.Equal(t, "mockAccessToken", responseRecorder.Header().Get("Access-Token"))
	})

	t.Run("LoginUser_InvalidRequest", func(t *testing.T) {
		loginPayload := &LoginPayload{
			Username: "",
			Password: "testpassword",
		}

		payload, _ := json.Marshal(loginPayload)
		request := httptest.NewRequest("POST", "/login", bytes.NewBuffer(payload))
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

		var response controllers.Response
		_ = json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.False(t, response.Success)
		assert.Contains(t, response.Error, "missing username")
	})
}
