package userscontroller

import (
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/julienschmidt/httprouter"
)

func (uc *UsersController) RemoveItem(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	itemID := ps.ByName("item_id")
	userID := r.URL.Query().Get("user_id")
	err := uc.userService.RemoveFromCart(userID, itemID)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusInternalServerError, err)
		return
	}

	controllers.RespondWithSuccess(rw, &controllers.Response{Success: true})
}
