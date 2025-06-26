package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	categories := r.Group("/product")
	{
		categories.GET("/:slug", controllers.GetProduct)
	}
}
