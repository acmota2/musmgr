package model

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	FilePath    string `gorm:"size:256"`
	Name        string
	Description string
}
