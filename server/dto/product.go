package dto

type ProductCharacteristic struct {
	FilterID   uint   `json:"filter_id"`
	FilterName string `json:"filter_name"`
	ValueID    uint   `json:"value_id"`
	Value      string `json:"value"`
	SeoSlug    string `json:"seo_slug,omitempty"`
	HasSeo     bool   `json:"has_seo"`
}

type ProductCategoryInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
