package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/akashgupta05/shopping-cart-go/app/controllers/adminscontroller"
	"github.com/akashgupta05/shopping-cart-go/app/controllers/authcontroller"
	"github.com/akashgupta05/shopping-cart-go/app/controllers/userscontroller"
	"github.com/akashgupta05/shopping-cart-go/app/models"
	"github.com/akashgupta05/shopping-cart-go/app/services"
	"github.com/julienschmidt/httprouter"
)

var authService *services.AuthService

func Init(router *httprouter.Router) {
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, "{ \"message\":\"Hello world!. I am Shopping cart Service.\",\"success\":true }")
	})
	router.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(404)
		fmt.Fprint(rw, "{ \"message\":\"Not Found.\",\"success\":true }")
	})
	authService = services.NewAuthService()

	authController := authcontroller.NewAuthController(authService)
	router.POST("/login", serveEndpoint(authController.LoginUser))

	userController := userscontroller.NewUsersController()
	router.POST("/register", serveEndpoint(userController.RegisterUser))

	router.GET("/items", userAuthMiddleware(userController.ListItems))
	router.GET("/users/cart", userAuthMiddleware(userController.GetCart))
	router.POST("/users/cart/add", userAuthMiddleware(userController.AddToCart))
	router.POST("/users/cart/remove", userAuthMiddleware(userController.RemoveItem))

	adminsController := adminscontroller.NewAdminsController()
	router.POST("/admins/suspend_user", adminAuthMiddleware(adminsController.SuspendUser))
	router.POST("/admins/add_items", adminAuthMiddleware(adminsController.AddItems))

}

func serveEndpoint(nextHandler func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params)) httprouter.Handle {
	return func(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		defer recoverFunc(w)
		setCommonHeaders(w)
		nextHandler(w, request, ps)
	}
}

func adminAuthMiddleware(nextHandler func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params)) httprouter.Handle {
	return func(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		defer recoverFunc(w)

		accessToken := request.Header.Get("Access-Token")
		valid, _ := authService.ValidateAccessToken(accessToken, string(models.ADMIN))
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		setCommonHeaders(w)
		nextHandler(w, request, ps)
	}
}

func userAuthMiddleware(nextHandler func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params)) httprouter.Handle {
	return func(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		defer recoverFunc(w)

		accessToken := request.Header.Get("Access-Token")
		valid, userID := authService.ValidateAccessToken(accessToken, string(models.USER))
		if !valid || userID == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		request.URL.RawQuery = strings.Join([]string{request.URL.RawQuery, fmt.Sprintf("user_id=%s", userID)}, "&")
		setCommonHeaders(w)
		nextHandler(w, request, ps)
	}
}

func recoverFunc(w http.ResponseWriter) {
	if recvr := recover(); recvr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("%v", recvr),
		}
		respBuytes, _ := json.Marshal(resp)
		setCommonHeaders(w)
		w.Write(respBuytes)
	}
}

func setCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization, AccessToken")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
}
