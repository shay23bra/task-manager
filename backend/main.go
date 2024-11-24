package main

import (
	"task-manager/routes"
	"task-manager/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	db := utils.InitDB()

	routes.RegisterTaskRoutes(r, db)

	r.Run(":8080")
}
