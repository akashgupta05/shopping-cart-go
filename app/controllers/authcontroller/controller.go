package authcontroller

import (
	"net/http"

	"github.com/akashgupta05/shopping-cart-go/app/services"
	"github.com/julienschmidt/httprouter"
)

type AuthControllerInterface interface {
	LoginUser(rw http.ResponseWriter, r *http.Request, ps httprouter.Params)
	LogoutUser(rw http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type AuthController struct {
	authService services.AuthServiceInterface
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
	}
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
