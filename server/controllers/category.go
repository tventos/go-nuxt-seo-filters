package controllers

import (
	"net/http"
	"server/config"
	"server/dto"
	"server/models"
	"server/repositories"
	"server/services"

	"github.com/gin-gonic/gin"
)

func GetAllCategoriesTree(c *gin.Context) {
	var categories []models.Category
	if err := config.DB.Order("parent_id NULLS FIRST, id").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch categories"})
		return
	}

	categoryMap := make(map[uint][]models.Category)
	rootCategories := []models.Category{}

	for _, cat := range categories {
		if cat.ParentID == nil {
			rootCategories = append(rootCategories, cat)
		} else {
			categoryMap[*cat.ParentID] = append(categoryMap[*cat.ParentID], cat)
		}
	}

	var buildTree func(cat models.Category) map[string]interface{}
	buildTree = func(cat models.Category) map[string]interface{} {
		node := map[string]interface{}{
			"id":   cat.ID,
			"name": cat.Name,
			"slug": cat.Slug,
		}

		children := []map[string]interface{}{}
		for _, child := range categoryMap[cat.ID] {
			children = append(children, buildTree(child))
		}

		if len(children) > 0 {
			node["children"] = children
		}

		return node
	}

	var result []map[string]interface{}
	for _, root := range rootCategories {
		result = append(result, buildTree(root))
	}

	c.JSON(http.StatusOK, result)
}

func GetProductsWithFilters(c *gin.Context) {
	path := c.Param("path")

	category, err := repositories.GetCategoryFromPath(path)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	pathFilterIDs, err := repositories.GetFilterValueIDsFromPath(path, category.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	queryFilterIDs := services.GetFilterValueIDsFromQuery(c, category)

	products, err := repositories.FindProducts(category.ID, pathFilterIDs, queryFilterIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query products"})
		return
	}

	result := services.MapProductsToDTO(products)
	c.JSON(http.StatusOK, result)
}

func GetFiltersByCategory(c *gin.Context) {
	categorySlug := c.Param("slug")

	filtersResp, err := services.GetFiltersByCategorySlug(categorySlug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, filtersResp)
}

func ApplyFilters(c *gin.Context) {
	var req dto.ApplyFiltersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := services.BuildFilterLink(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetCategoryInfo(c *gin.Context) {
	path := c.Param("path")

	category, err := repositories.GetCategoryFromPath(path)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	pathFilterIDs, err := repositories.GetFilterValueIDsFromPath(path, category.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if len(pathFilterIDs) > 0 {
		seoFilterValue, err := repositories.GetFilterSeoValue(pathFilterIDs[0], category.ID)

		if err == nil {
			c.JSON(http.StatusOK, dto.CategoryInfoResponse{
				Title:         seoFilterValue.Title,
				Description:   seoFilterValue.Description,
				H1:            seoFilterValue.H1,
				Content:       seoFilterValue.Content,
				FilterValueId: seoFilterValue.FilterValueID,
				FilterSlug:    seoFilterValue.FilterValue.Filter.Slug,
			})
			return
		}
	}

	c.JSON(http.StatusOK, dto.CategoryInfoResponse{
		Title:       category.Title,
		Description: category.Description,
		H1:          category.H1,
		Content:     category.Content,
	})
}
