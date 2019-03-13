package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"varchar(64)" json:"user_name"`
	Passowrd string `gorm:"varchar(64)" json:"password"`
	Email    string `gorm:"varchar(128)" json:"email"`
	Role     int    `json:"role"`
}
