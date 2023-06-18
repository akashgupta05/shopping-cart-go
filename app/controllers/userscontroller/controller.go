package userscontroller

import (
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/services"
	"github.com/julienschmidt/httprouter"
)

type UsersControllerInterface interface {
	RegisterUser(rw http.ResponseWriter, r *http.Request, ps httprouter.Params)
	ListItems(rw http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddToCart(rw http.ResponseWriter, r *http.Request, ps httprouter.Params)
	RemoveItem(rw http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetCart(rw http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type UsersController struct {
	userService services.UserServiceInterface
	itemService services.ItemServiceInterface
}

func NewUsersController() *UsersController {
	return &UsersController{
		userService: services.NewUserService(),
		itemService: services.NewItemService(),
	}
}

type RegisterUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AddItemPayload struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}

type RemoveItemPayload struct {
	ItemID string `json:"item_id"`
}
