package userscontroller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/julienschmidt/httprouter"
)

func (uc *UsersController) AddToCart(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bodyBytes, err := controllers.ReadBodyBytes(r)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	addItemPayload := &AddItemPayload{}
	err = json.Unmarshal(bodyBytes, addItemPayload)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	if err = validateAddToCartRequest(addItemPayload); err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	userID := r.URL.Query().Get("user_id")
	err = uc.userService.AddToCart(userID, addItemPayload.ItemID, addItemPayload.Quantity)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusInternalServerError, err)
		return
	}

	controllers.RespondWithSuccess(rw, &controllers.Response{Success: true})
}

func validateAddToCartRequest(addToCart *AddItemPayload) error {
	if addToCart.ItemID == "" {
		return errors.New("missing item_id")
	}

	if addToCart.Quantity <= 0 {
		return errors.New("missing quantity")
	}

	return nil
}
