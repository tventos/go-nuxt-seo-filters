package models

type Filter struct {
	ID     uint
	Name   string
	Slug   string
	Type   string
	Values []FilterValue
}

type FilterValue struct {
	ID       uint   `gorm:"primaryKey"`
	Value    string `gorm:"not null"`
	FilterID uint   `gorm:"index"`
	Filter   Filter
	Products []Product       `gorm:"many2many:product_filter_values;"`
	Seo      *SeoFilterValue `gorm:"foreignKey:FilterValueID"`
}

type SeoFilterValue struct {
	ID            uint
	FilterValueID uint
	CategoryId    uint
	Title         string
	Description   string
	H1            string
	Content       string
	Slug          string
	FilterValue   FilterValue `gorm:"foreignKey:FilterValueID"`
}
