package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/akashgupta05/shopping-cart-go/app/routes"
	"github.com/akashgupta05/shopping-cart-go/config/db"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// func main() {
// 	router := mux.NewRouter()

// 	router.HandleFunc("/register", handlers.Register).Methods("POST")
// 	router.HandleFunc("/login", handlers.Login).Methods("POST")

// 	// Restricted routes for admin
// 	router.HandleFunc("/admin/add-item", utils.RBAC(utils.RoleAdmin, handlers.AddItem)).Methods("POST")
// 	router.HandleFunc("/admin/suspend-user", utils.RBAC(utils.RoleAdmin, handlers.SuspendUser)).Methods("POST")

// 	// Restricted routes for user
// 	router.HandleFunc("/user/list-items", utils.RBAC(utils.RoleUser, handlers.ListItems)).Methods("GET")
// 	router.HandleFunc("/user/add-to-cart", utils.RBAC(utils.RoleUser, handlers.AddToCart)).Methods("POST")
// 	router.HandleFunc("/user/remove-from-cart", utils.RBAC(utils.RoleUser, handlers.RemoveFromCart)).Methods("POST")

// 	log.Fatal(http.ListenAndServe(":8000", router))
// }

var router = httprouter.New()

func init() {
	setupDB()
	routes.Init(router)
	log.Info("Routes initialized")
}

func main() {
	defer db.Close()

	var port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Debug("Starting server at port", port)

	server := &http.Server{Addr: port, Handler: router}

	go server.ListenAndServe()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	<-ch

	log.Debug("Stopping server at port", port)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Failed to shut down the server gracefully")
	}
}

func setupDB() {
	if err := db.Connect(os.Getenv("DATABASE_URL")); err != nil {
		panic(err)
	}
	log.Info("Database connection : Done")
}
