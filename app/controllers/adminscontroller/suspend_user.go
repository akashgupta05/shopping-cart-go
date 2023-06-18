package adminscontroller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/julienschmidt/httprouter"
)

func (ac *AdminsController) SuspendUser(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bodyBytes, err := controllers.ReadBodyBytes(r)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	suspendUserPayload := &SuspendUserPayload{}
	err = json.Unmarshal(bodyBytes, suspendUserPayload)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	if err = validateSuspendUserRequest(suspendUserPayload); err != nil {
		controllers.RespondWithError(rw, http.StatusBadRequest, err)
		return
	}

	err = ac.adminService.SuspendUser(suspendUserPayload.UserID)
	if err != nil {
		controllers.RespondWithError(rw, http.StatusInternalServerError, err)
		return
	}

	controllers.RespondWithSuccess(rw, &controllers.Response{Success: true})
}

func validateSuspendUserRequest(suspendUserPayload *SuspendUserPayload) error {
	if suspendUserPayload.UserID == "" {
		return errors.New("missing user_id")
	}

	return nil
}
