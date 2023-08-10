package main

import (
	"log"

	"github.com/bluebasooo/tic-tac-toe_backend/db"
	"github.com/bluebasooo/tic-tac-toe_backend/internal/user"
	"github.com/bluebasooo/tic-tac-toe_backend/internal/ws"
	"github.com/bluebasooo/tic-tac-toe_backend/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize db connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}
