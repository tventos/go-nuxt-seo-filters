package repositories

import (
	"fmt"
	"server/config"
	"server/models"
)

func FindProducts(categoryID uint, pathFilters []uint, queryFilters map[uint][]uint) ([]models.Product, error) {
	query := config.DB.Model(&models.Product{}).
		Joins("JOIN product_categories pc ON pc.product_id = products.id").
		Where("pc.category_id = ?", categoryID).
		Preload("Categories").
		Preload("FilterValues.Filter")

	for i, fvID := range pathFilters {
		alias := fmt.Sprintf("pfv_path_%d", i)
		query = query.Joins(fmt.Sprintf("JOIN product_filter_values %s ON %s.product_id = products.id", alias, alias)).
			Where(fmt.Sprintf("%s.filter_value_id = ?", alias), fvID)
	}

	i := 0
	for _, ids := range queryFilters {
		if len(ids) == 0 {
			continue
		}
		alias := fmt.Sprintf("pfv_query_%d", i)
		query = query.Joins(fmt.Sprintf("JOIN product_filter_values %s ON %s.product_id = products.id", alias, alias)).
			Where(fmt.Sprintf("%s.filter_value_id IN ?", alias), ids)
		i++
	}

	var products []models.Product
	err := query.Find(&products).Error
	return products, err
}
