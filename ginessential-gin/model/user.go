package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Tel  string `gorm:"type:varchar(11);not null;unique"`
	Pwd  string `gorm:"size:255,not null"`
}
