package main

import (
	"server/config"
	"server/middlewares"
	"server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	config.ConnectDB()

	routes.RegisterCategoryRoutes(r)
	routes.RegisterProductRoutes(r)

	r.Run(":8080")

}
