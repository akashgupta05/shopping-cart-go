package userscontroller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/akashgupta05/shopping-cart-go/app/controllers"
	"github.com/julienschmidt/httprouter"
)

var page = "page"
var perpage = "per_page"

func (uc *UsersController) ListItems(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	pageNo := queryValues.Get(page)
	perPage := queryValues.Get(perpage)

	pageNumber, err := strconv.Atoi(pageNo)
	if err != nil {
		pageNumber = 1
	}

	pageSize, err := strconv.Atoi(perPage)
	if err != nil {
		pageSize = 10
	}

	if pageNumber <= 0 || pageSize <= 0 {
		controllers.RespondWithError(rw, http.StatusBadRequest, errors.New("invalid page or per_page"))
		return
	}

	items, err := uc.itemService.ListItems()
	if err != nil {
		controllers.RespondWithError(rw, http.StatusInternalServerError, errors.New("Failed to fetch items"))
		return
	}

	controllers.RespondWithSuccess(rw, &controllers.Response{Success: true, Data: items})
}
