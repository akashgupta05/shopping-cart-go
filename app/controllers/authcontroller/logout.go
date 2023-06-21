package authcontroller

import (
	"net/http"
	"time"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/julienschmidt/httprouter"
)

func (ac *AuthController) LogoutUser(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})

	controllers.RespondWithSuccess(rw, &controllers.Response{Success: true})
}
