package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(r *gin.Engine) {
	categories := r.Group("/category")
	{
		categories.GET("", controllers.GetAllCategoriesTree)
		categories.GET("/info/*path", controllers.GetCategoryInfo)
		categories.GET("/filters/list/:slug", controllers.GetFiltersByCategory)
		categories.POST("/filters/apply", controllers.ApplyFilters)
		categories.GET("/products/*path", controllers.GetProductsWithFilters)
	}
}
