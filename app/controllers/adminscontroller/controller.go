package adminscontroller

import (
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/services"
	"github.com/julienschmidt/httprouter"
)

type AdminsControllerInterface interface {
	SuspendUser(rw http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddItems(rw http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type AdminsController struct {
	adminService services.AdminServiceInterface
	itemService  services.ItemServiceInterface
}

func NewAdminsController() *AdminsController {
	return &AdminsController{
		adminService: services.NewAdminService(),
		itemService:  services.NewItemService(),
	}
}

type AddItemsPayload struct {
	ItemName string `json:"item_name"`
	Quantity int    `json:"quantity"`
}

type SuspendUserPayload struct {
	UserID string `json:"user_id"`
}
