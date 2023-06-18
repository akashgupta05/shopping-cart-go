package authcontroller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/julienschmidt/httprouter"
)

func (ac *AuthController) LoginUser(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bodyBytes, err := controllers.ReadBodyBytes(r)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	loginPayload := &LoginPayload{}
	err = json.Unmarshal(bodyBytes, loginPayload)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	if err = validateLoginRequest(loginPayload); err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	accessToken, err := ac.authService.Login(loginPayload.Username, loginPayload.Password)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusInternalServerError, err)
		return
	}

	controllers.RespondWithAccessTokenSuccess(rw, accessToken, &controllers.Response{Success: true})
}

func validateLoginRequest(loginPayload *LoginPayload) error {
	if loginPayload.Username == "" {
		return errors.New("missing user_id")
	}

	if loginPayload.Password == "" {
		return errors.New("missing password")
	}

	return nil
}
