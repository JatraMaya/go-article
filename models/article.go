package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string
	Author  string
	Content string `gorm:"type:text"`
}
