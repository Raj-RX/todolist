package main

import (
	"fmt"
	"github.com/tejashwikalptaru/tutorial/database"
	"github.com/tejashwikalptaru/tutorial/server"
)

func main() {
	err := database.ConnectAndMigrate("localhost", "5433", "tutorial", "local", "local", database.SSLModeDisable)
	if err != nil {
		fmt.Println("check-->")
		panic(err)
	}

	fmt.Println("connected")
	srv := server.SetupRoutes()
	err = srv.Run(":8080")
	if err != nil {
		panic(err)
	}

}
