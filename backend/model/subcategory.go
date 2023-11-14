package model

import (
	database "backend/db"

	"gorm.io/gorm"
)

type SubCategory struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  uint64 `json:"category_id"`
}

func (s *SubCategory) Save() (*SubCategory, error) {
	err := database.PsqlDB.Select("name", "category_id").Create(s).Error
	if err != nil {
		return &SubCategory{}, err
	}
	return s, nil
}
