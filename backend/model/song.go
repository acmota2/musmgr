package model

import (
	database "backend/db"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Name        string
	Description string
	Tonality    string
	Files       []*SongFile
	Events      []*Event
	SubCategory []*SubCategory
}

func (s *Song) Save() (*Song, error) {
	err := database.PsqlDB.Select("name", "description", "tonality").Create(s).Error
	// provavelmente tenho que pôr algo aqui para adicionar à song_subcategory
	// mas depois vejo como fazer isso
	if err != nil {
		return &Song{}, err
	}
	return s, nil
}
