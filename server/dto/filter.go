package dto

type ApplyFiltersRequest struct {
	CategorySlug string              `json:"category_slug"`
	Filters      map[string][]uint64 `json:"filters"`
}

type ApplyFiltersResponse struct {
	Link string `json:"link"`
}
