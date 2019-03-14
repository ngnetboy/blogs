package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	gorm.Model

	Name     string      `gorm:"varchar(128)" json:"name"`
	Content  string      `gorm:"text" json:"content"`
	Summary  string      `gorm:"text" json:"summary"`
	UserID   uint        `json:"user_id"`
	Amount   uint        `json:"amount"`
	Draft    int         `json:"draft"`
	Category []*Category `gorm:"many2many:article_category" json:"category"`
	Tags     []*Tag      `gorm:"many2many:article_tag" json:"tag"`
}

type ArticleBase struct {
	ID       uint      `json:"ID"`
	CreateAt time.Time `json:"CreateTime"`
	Name     string    `json:"name"`
	Summary  string    `json:"summary"`
	Amount   uint      `json:"amount"`
}
