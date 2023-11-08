package model

import (
	database "backend/db"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *Category) Save() (*Category, error) {
	err := database.DB.Create(c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}
