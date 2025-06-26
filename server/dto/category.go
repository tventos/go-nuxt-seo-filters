package dto

type CategoryProductsResponse struct {
	ID           uint                          `json:"id"`
	Name         string                        `json:"name"`
	Description  string                        `json:"description"`
	Slug         string                        `json:"slug"`
	Categories   []CategoryResponse            `json:"categories"`
	FilterValues []CategoryFilterValueResponse `json:"filter_values"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CategoryFilterValueResponse struct {
	ID     uint                        `json:"id"`
	Value  string                      `json:"value"`
	Filter CategoryFilterShortResponse `json:"filter"`
}

type CategoryFilterShortResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CategoryFilterValueShortResponse struct {
	ID    uint   `json:"id"`
	Value string `json:"value"`
}

type CategoryFilterResponse struct {
	ID     uint                               `json:"id"`
	Name   string                             `json:"name"`
	Type   string                             `json:"type"`
	Slug   string                             `json:"slug"`
	Values []CategoryFilterValueShortResponse `json:"values"`
}

type CategoryInfoResponse struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	H1            string `json:"h1"`
	Content       string `json:"content"`
	FilterValueId uint   `json:"filter_value_id"`
	FilterSlug    string `json:"filter_slug"`
}
