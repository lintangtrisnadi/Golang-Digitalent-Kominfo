package main

import (
	"restapigo/database"
	"restapigo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
