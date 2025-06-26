package repositories

import (
	"errors"
	"fmt"
	"server/config"
	"server/models"
	"strings"
)

func GetFilterValueIDsFromPath(path string, categoryID uint) ([]uint, error) {
	segments := strings.Split(strings.Trim(path, "/"), "/")[1:]
	var filterValueIDs []uint

	for _, seoSlug := range segments {
		var sfv models.SeoFilterValue
		err := config.DB.
			Joins("JOIN filter_values fv ON fv.id = seo_filter_values.filter_value_id").
			Joins("JOIN filters f ON f.id = fv.filter_id").
			Joins("JOIN category_filters cf ON cf.filter_id = f.id").
			Where("seo_filter_values.slug = ? AND cf.category_id = ?", seoSlug, categoryID).
			First(&sfv).Error

		if err != nil {
			return nil, fmt.Errorf("seo filter value not found: %s", seoSlug)
		}

		filterValueIDs = append(filterValueIDs, sfv.FilterValueID)
	}

	return filterValueIDs, nil
}

func GetFilterValueByID(id uint64) (models.FilterValue, error) {
	var fv models.FilterValue
	err := config.DB.Preload("Seo").Where("id = ?", id).First(&fv).Error
	return fv, err
}

func GetFilterSeoValue(filterValueId uint, categoryID uint) (models.SeoFilterValue, error) {
	var fsv models.SeoFilterValue
	err := config.DB.
		Preload("FilterValue.Filter").
		Where("filter_value_id = ? AND category_id = ?", filterValueId, categoryID).First(&fsv).Error

	if err != nil {
		return models.SeoFilterValue{}, errors.New("seo filter not found")
	}

	return fsv, err
}
