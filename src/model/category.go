package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model

	Name        string     `gorm:"type:varchar(64)" json:"name"`
	Description string     `gorm:"type:varchar(512)" json:"description"`
	Article     []*Article `gorm:"many2many:article_category" json:"article"`
}

type CategoryCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
