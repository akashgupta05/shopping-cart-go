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
