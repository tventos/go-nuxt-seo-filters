package controllers

import (
	"net/http"
	"server/config"
	"server/dto"
	"server/models"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	slug := c.Param("slug")

	var product models.Product
	err := config.DB.
		Preload("FilterValues.Filter").
		Preload("FilterValues.Seo").
		Preload("Categories").
		Where("slug = ?", slug).
		First(&product).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	var characteristics []dto.ProductCharacteristic
	for _, fv := range product.FilterValues {
		ch := dto.ProductCharacteristic{
			FilterID:   fv.Filter.ID,
			FilterName: fv.Filter.Name,
			ValueID:    fv.ID,
			Value:      fv.Value,
			HasSeo:     fv.Seo != nil,
		}
		if fv.Seo != nil {
			ch.SeoSlug = fv.Seo.Slug
		}
		characteristics = append(characteristics, ch)
	}

	var categoryInfo *dto.ProductCategoryInfo
	if len(product.Categories) > 0 {
		cat := product.Categories[len(product.Categories)-1]
		categoryInfo = &dto.ProductCategoryInfo{
			ID:   cat.ID,
			Name: cat.Name,
			Slug: cat.Slug,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"id":              product.ID,
		"name":            product.Name,
		"description":     product.Description,
		"slug":            product.Slug,
		"category":        categoryInfo,
		"characteristics": characteristics,
	})
}
