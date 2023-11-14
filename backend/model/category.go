package model

import (
	database "backend/db"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	SubCategories []SubCategory `json:"subcategories"`
}

func GetAllCategories() (*[]Category, error) {
	var categories *[]Category
	err := database.PsqlDB.Find(&categories).Error

	if err != nil {
		return &[]Category{}, err
	}
	return categories, nil
}
