package models

type Category struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Slug        string `gorm:"uniqueIndex;not null"`
	H1          string
	Title       string
	Description string
	Content     string
	ParentID    *uint
	Parent      *Category
	Children    []Category `gorm:"foreignKey:ParentID"`
	Products    []Product  `gorm:"many2many:product_categories;"`
	Filters     []Filter   `gorm:"many2many:category_filters"`
}
