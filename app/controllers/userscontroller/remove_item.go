package userscontroller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/julienschmidt/httprouter"
)

func (uc *UsersController) RemoveItem(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bodyBytes, err := controllers.ReadBodyBytes(r)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	removeItemPayload := &RemoveItemPayload{}
	err = json.Unmarshal(bodyBytes, removeItemPayload)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	if err = validateRemoveItemRequest(removeItemPayload); err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	userID := r.URL.Query().Get("user_id")
	err = uc.userService.RemoveFromCart(userID, removeItemPayload.ItemID)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusInternalServerError, err)
		return
	}

	controllers.RespondWithSuccess(rw, &controllers.Response{Success: true})
}

func validateRemoveItemRequest(removeItemPayload *RemoveItemPayload) error {
	if removeItemPayload.ItemID == "" {
		return errors.New("missing item_id")
	}

	return nil
}
