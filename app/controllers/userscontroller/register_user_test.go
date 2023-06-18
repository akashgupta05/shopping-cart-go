package userscontroller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestUsersController_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock.NewMockUserServiceInterface(ctrl)

	usersController := NewUsersController()
	usersController.userService = mockUserService

	router := httprouter.New()
	router.POST("/register", usersController.RegisterUser)

	t.Run("RegisterUser_Success", func(t *testing.T) {
		mockUsername := "john_doe"
		mockPassword := "password"
		mockRole := "user"

		mockUser := &models.User{
			ID:       "123",
			Username: mockUsername,
			Role:     models.Role(mockRole),
		}

		mockUserService.EXPECT().RegisterUser(mockUsername, mockPassword, models.Role(mockRole)).Return(mockUser, nil)

		payload := `{"username":"john_doe","password":"password","role":"user"}`
		request := httptest.NewRequest("POST", "/register", strings.NewReader(payload))
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)

		var response controllers.Response
		err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.Empty(t, response.Error)

		var userData models.User
		respBytes, _ := json.Marshal(response.Data)
		err = json.Unmarshal(respBytes, &userData)

		assert.NoError(t, err)
		assert.Equal(t, mockUser.ID, userData.ID)
		assert.Equal(t, mockUser.Username, userData.Username)
		assert.Equal(t, mockUser.Role, userData.Role)
	})

	t.Run("RegisterUser_InvalidRequest", func(t *testing.T) {
		payload := `{"username":"","password":"password","role":"user"}` // Invalid payload with missing username

		request := httptest.NewRequest("POST", "/register", strings.NewReader(payload))
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

		var response controllers.Response
		err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.False(t, response.Success)
		assert.Contains(t, response.Error, "missing username")
	})

	t.Run("RegisterUser_Error", func(t *testing.T) {
		mockUsername := "john_doe"
		mockPassword := "password"
		mockRole := "user"

		mockUserService.EXPECT().RegisterUser(mockUsername, mockPassword, models.Role(mockRole)).Return(nil, errors.New("failed to register user"))

		payload := `{"username":"john_doe","password":"password","role":"user"}`
		request := httptest.NewRequest("POST", "/register", strings.NewReader(payload))
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

		var response controllers.Response
		err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.False(t, response.Success)
		assert.Contains(t, response.Error, "failed to register user")
	})
}
