package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model

	Name    string     `gorm:"varchar(64)" json:"name"`
	Article []*Article `gorm:"many2many:article_tag" json:"article"`
}
