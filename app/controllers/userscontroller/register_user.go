package userscontroller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (uc *UsersController) RegisterUser(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bodyBytes, err := controllers.ReadBodyBytes(r)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	registerUser := &RegisterUserPayload{}
	err = json.Unmarshal(bodyBytes, registerUser)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	if err = validateRegisterRequest(registerUser); err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	user, err := uc.userService.RegisterUser(registerUser.Username, registerUser.Password, models.Role(registerUser.Role))
	if err != nil {
		log.Warnf("Error while registering user : %v", err)
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	controllers.RespondWithSuccess(rw, &controllers.Response{Success: true, Data: user})
}

func validateRegisterRequest(registerUser *RegisterUserPayload) error {
	if registerUser.Username == "" {
		return errors.New("missing username")
	}

	if registerUser.Password == "" {
		return errors.New("missing password")
	}

	if registerUser.Role == "" {
		return errors.New("missing role")
	}

	if registerUser.Role != string(models.ADMIN) && registerUser.Role != string(models.USER) {
		return errors.New("invalid role")
	}

	return nil
}
