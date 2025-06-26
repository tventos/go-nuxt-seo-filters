package services

import (
	"errors"
	"fmt"
	"server/dto"
	"server/repositories"
	"strings"
)

func GetFiltersByCategorySlug(slug string) ([]dto.CategoryFilterResponse, error) {
	category, err := repositories.GetCategoryWithFilters(slug)
	if err != nil {
		return nil, errors.New("category not found")
	}

	var filtersResp []dto.CategoryFilterResponse
	for _, filter := range category.Filters {
		var valuesResp []dto.CategoryFilterValueShortResponse
		for _, val := range filter.Values {
			valuesResp = append(valuesResp, dto.CategoryFilterValueShortResponse{
				ID:    val.ID,
				Value: val.Value,
			})
		}

		filtersResp = append(filtersResp, dto.CategoryFilterResponse{
			ID:     filter.ID,
			Name:   filter.Name,
			Type:   filter.Type,
			Values: valuesResp,
			Slug:   filter.Slug,
		})
	}

	return filtersResp, nil
}

func BuildFilterLink(req dto.ApplyFiltersRequest) (dto.ApplyFiltersResponse, error) {
	_, err := repositories.GetCategoryBySlug(req.CategorySlug)
	if err != nil {
		return dto.ApplyFiltersResponse{}, fmt.Errorf("category not found")
	}

	segments := []string{req.CategorySlug}
	queryParams := map[string][]string{}
	addedSeoFilter := ""

	for filterSlug, ids := range req.Filters {
		if len(ids) != 1 {
			for _, id := range ids {
				queryParams[filterSlug] = append(queryParams[filterSlug], fmt.Sprintf("%d", id))
			}
			continue
		}

		id := ids[0]
		fv, err := repositories.GetFilterValueByID(id)
		if err != nil {
			continue
		}

		if addedSeoFilter == "" && fv.Seo != nil && fv.Seo.Slug != "" {
			segments = append(segments, fv.Seo.Slug)
			addedSeoFilter = filterSlug
		} else {
			queryParams[filterSlug] = append(queryParams[filterSlug], fmt.Sprintf("%d", id))
		}
	}

	link := "/category/" + strings.Join(segments, "/")

	if len(queryParams) > 0 {
		var qp []string
		for key, values := range queryParams {
			qp = append(qp, fmt.Sprintf("%s=%s", key, strings.Join(values, ",")))
		}
		link += "?" + strings.Join(qp, "&")
	}

	return dto.ApplyFiltersResponse{Link: link}, nil
}
