package models

type Product struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Description  string
	Slug         string        `gorm:"uniqueIndex;not null"`
	Categories   []Category    `gorm:"many2many:product_categories;"`
	FilterValues []FilterValue `gorm:"many2many:product_filter_values;"`
}
