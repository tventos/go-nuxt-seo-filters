package services

import (
	"server/dto"
	"server/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetFilterValueIDsFromQuery(c *gin.Context, category models.Category) map[uint][]uint {
	result := make(map[uint][]uint)

	for _, filter := range category.Filters {
		param := c.Query(filter.Slug)
		if param == "" {
			continue
		}

		ids := strings.Split(param, ",")
		for _, idStr := range ids {
			id, err := strconv.ParseUint(idStr, 10, 64)
			if err == nil {
				result[filter.ID] = append(result[filter.ID], uint(id))
			}
		}
	}

	return result
}

func MapProductsToDTO(products []models.Product) []dto.CategoryProductsResponse {
	var result []dto.CategoryProductsResponse

	for _, p := range products {
		var cats []dto.CategoryResponse
		for _, c := range p.Categories {
			cats = append(cats, dto.CategoryResponse{
				ID:   c.ID,
				Name: c.Name,
				Slug: c.Slug,
			})
		}

		var vals []dto.CategoryFilterValueResponse
		for _, fv := range p.FilterValues {
			vals = append(vals, dto.CategoryFilterValueResponse{
				ID:    fv.ID,
				Value: fv.Value,
				Filter: dto.CategoryFilterShortResponse{
					ID:   fv.Filter.ID,
					Name: fv.Filter.Name,
				},
			})
		}

		result = append(result, dto.CategoryProductsResponse{
			ID:           p.ID,
			Name:         p.Name,
			Description:  p.Description,
			Slug:         p.Slug,
			Categories:   cats,
			FilterValues: vals,
		})
	}

	if result == nil {
		result = make([]dto.CategoryProductsResponse, 0)
	}

	return result
}
