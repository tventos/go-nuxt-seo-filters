package repositories

import (
	"errors"
	"server/config"
	"server/models"
	"strings"
)

func GetCategoryFromPath(path string) (models.Category, error) {
	segments := strings.Split(strings.Trim(path, "/"), "/")
	if len(segments) < 1 {
		return models.Category{}, errors.New("invalid path")
	}
	categorySlug := segments[0]

	var category models.Category
	err := config.DB.Preload("Filters.Values").
		Where("slug = ?", categorySlug).
		First(&category).Error

	if err != nil {
		return models.Category{}, errors.New("category not found")
	}

	return category, nil
}

func GetCategoryWithFilters(slug string) (models.Category, error) {
	var category models.Category
	err := config.DB.Preload("Filters.Values").Where("slug = ?", slug).First(&category).Error
	return category, err
}

func GetCategoryBySlug(slug string) (models.Category, error) {
	var category models.Category
	err := config.DB.Where("slug = ?", slug).First(&category).Error
	return category, err
}
