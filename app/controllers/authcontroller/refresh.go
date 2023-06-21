package authcontroller

import (
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/julienschmidt/httprouter"
)

func (ac *AuthController) Refresh(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.URL.Query().Get("user_id")
	jwtToken, expiresAt, err := ac.authService.RefreshJWT(userID)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusInternalServerError, err)
		return
	}

	controllers.RespondWithJWTSuccess(rw, jwtToken, expiresAt, &controllers.Response{Success: true})
}
