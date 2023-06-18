package userscontroller

import (
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/julienschmidt/httprouter"
)

func (uc *UsersController) GetCart(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.URL.Query().Get("user_id")
	cartItems, err := uc.userService.GetCartItems(userID)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusInternalServerError, err)
		return
	}

	controllers.RespondWithSuccess(rw, &controllers.Response{Success: true, Data: cartItems})
}
