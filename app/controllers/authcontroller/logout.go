package authcontroller

import (
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/julienschmidt/httprouter"
)

func (ac *AuthController) LogoutUser(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	accessToken := r.Header.Get("Access-Token")
	if err := ac.authService.Logout(accessToken); err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	controllers.RespondWithSuccess(rw, &controllers.Response{Success: true})
}
