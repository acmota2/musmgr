package model

import (
	database "backend/db"
	songs "backend/songs"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Name        string
	Description string
	Tonality    songs.Tonality
}

func (s *Song) Save() (*Song, error) {
	err := database.DB.Create(s).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}
