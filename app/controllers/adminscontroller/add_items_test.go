package adminscontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"

	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type ItemsMatcher struct {
	Items []*models.Item
}

func (m *ItemsMatcher) Matches(x interface{}) bool {
	items, ok := x.([]*models.Item)
	if !ok {
		return false
	}

	if len(items) != len(m.Items) {
		return false
	}

	for i, item := range items {
		return item.Name == m.Items[i].Name
	}

	return true
}

func (m *ItemsMatcher) String() string {
	return fmt.Sprintf("matches items")
}

func TestAdminsController_AddItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminService := mock.NewMockAdminServiceInterface(ctrl)
	mockItemService := mock.NewMockItemServiceInterface(ctrl)

	adminsController := NewAdminsController()
	adminsController.adminService = mockAdminService
	adminsController.itemService = mockItemService

	router := httprouter.New()
	router.POST("/items", adminsController.AddItems)

	t.Run("AddItems_Success", func(t *testing.T) {
		itemPayload := []*AddItemsPayload{
			{ItemName: "Item 1", Quantity: 10},
			{ItemName: "Item 2", Quantity: 5},
		}

		items := []*models.Item{
			{Name: "Item 1", Quantity: 10},
			{Name: "Item 2", Quantity: 5},
		}

		itemsMatcher := &ItemsMatcher{Items: items}

		mockAdminService.EXPECT().AddItems(itemsMatcher).Return(items, nil)

		payload, _ := json.Marshal(itemPayload)
		request := httptest.NewRequest("POST", "/items", bytes.NewBuffer(payload))
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)

		var response controllers.Response
		_ = json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.True(t, response.Success)
		actualItems := []*models.Item{}
		respBytes, _ := json.Marshal(response.Data)
		_ = json.Unmarshal(respBytes, &actualItems)
		assert.Equal(t, len(items), len(actualItems))
	})

	t.Run("AddItems_InvalidRequest", func(t *testing.T) {
		itemPayload := []*AddItemsPayload{
			{ItemName: "", Quantity: 10},      // Invalid: missing item name
			{ItemName: "Item 2", Quantity: 0}, // Invalid: zero quantity
		}

		payload, _ := json.Marshal(itemPayload)
		request := httptest.NewRequest("POST", "/items", bytes.NewBuffer(payload))
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

		var response controllers.Response
		_ = json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		assert.False(t, response.Success)
		assert.Contains(t, response.Error, "missing item_name")
	})
}
