package userscontroller

import (
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/julienschmidt/httprouter"
)

func (uc *UsersController) ListItems(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	items, err := uc.itemService.ListItems()
	if err != nil {
		controllers.RespondWithError(rw, http.StatusInternalServerError, err)
		return
	}

	controllers.RespondWithSuccess(rw, &controllers.Response{Success: true, Data: items})
}
