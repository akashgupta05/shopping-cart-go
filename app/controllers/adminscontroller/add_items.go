package adminscontroller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/julienschmidt/httprouter"
)

func (ac *AdminsController) AddItems(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bodyBytes, err := controllers.ReadBodyBytes(r)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	addItemPayload := []*AddItemsPayload{}
	err = json.Unmarshal(bodyBytes, &addItemPayload)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	if err = validateAddItemsRequest(addItemPayload); err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	items := translateRequestPayloadToModel(addItemPayload)

	addedItems, err := ac.adminService.AddItems(items)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusInternalServerError, err)
		return
	}

	controllers.RespondWithSuccess(rw, &controllers.Response{Success: true, Data: addedItems})
}

func translateRequestPayloadToModel(addItems []*AddItemsPayload) []*models.Item {
	items := []*models.Item{}
	for _, addItem := range addItems {
		items = append(items, &models.Item{
			Quantity: addItem.Quantity, Name: addItem.ItemName,
		})
	}

	return items
}

func validateAddItemsRequest(addItems []*AddItemsPayload) error {
	for _, item := range addItems {
		if item.ItemName == "" {
			return errors.New("missing item_name")
		}

		if item.Quantity <= 0 {
			return errors.New("missing item_id")
		}
	}

	return nil
}
