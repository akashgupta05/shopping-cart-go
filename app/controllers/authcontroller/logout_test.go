package authcontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestAuthController_LogoutUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authController := NewAuthController()

	router := httprouter.New()
	router.POST("/logout", authController.LogoutUser)

	t.Run("LogoutUser_Success", func(t *testing.T) {
		request := httptest.NewRequest("POST", "/logout", nil)
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)

		cookies := responseRecorder.Result().Cookies()
		var tokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "token" {
				tokenCookie = cookie
				break
			}
		}

		// Assert that the "token" cookie has expired
		assert.NotNil(t, tokenCookie)
		assert.True(t, tokenCookie.Expires.Before(time.Now()))
	})
}
