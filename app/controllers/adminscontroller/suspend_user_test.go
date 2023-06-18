package adminscontroller

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

func TestAdminsController_SuspendUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminService := mock.NewMockAdminServiceInterface(ctrl)

	adminsController := NewAdminsController()
	adminsController.adminService = mockAdminService

	router := httprouter.New()
	router.POST("/admins/suspend_user", adminsController.SuspendUser)

	t.Run("SuspendUser_Success", func(t *testing.T) {
		suspendUserPayload := &SuspendUserPayload{UserID: "123"}

		mockAdminService.EXPECT().SuspendUser(suspendUserPayload.UserID).Return(nil)

		payload, _ := json.Marshal(suspendUserPayload)
		request := httptest.NewRequest("POST", "/admins/suspend_user", bytes.NewBuffer(payload))
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)

		var response controllers.Response
		_ = json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.True(t, response.Success)
		assert.Empty(t, response.Error)
	})

	t.Run("SuspendUser_InvalidRequest", func(t *testing.T) {
		suspendUserPayload := &SuspendUserPayload{UserID: ""}

		payload, _ := json.Marshal(suspendUserPayload)
		request := httptest.NewRequest("POST", "/admins/suspend_user", bytes.NewBuffer(payload))
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

		var response controllers.Response
		_ = json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.False(t, response.Success)
		assert.Contains(t, response.Error, "missing user_id")
	})
}
