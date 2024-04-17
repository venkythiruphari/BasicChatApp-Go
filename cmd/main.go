package main

import (
	"chat/db"
	"chat/internal/user"
	"chat/internal/ws"
	"chat/router"
	"fmt"
)

func main() {
	dbConn, err := db.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	fmt.Println("Hello Gunni DataBase Connected Successful")
	defer dbConn.Close()

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	fmt.Println("Hello Gunni Port Started On :8080 Successful")
	router.Start("0.0.0.0:8080") //router.Start(":3000")
}
