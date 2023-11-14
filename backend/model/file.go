package model

import (
	database "backend/db"

	"gorm.io/gorm"
)

type FileType uint8

const (
	text  = iota
	score = iota
)

type SongFile struct {
	gorm.Model
	Path string
	Type FileType
}

func (f *SongFile) Save() (*SongFile, error) {
	err := database.PsqlDB.Create(f).Error

	if err != nil {
		return &SongFile{}, err
	}
	return f, nil
}

// vais ter que ser mais inteligente a lidar com a abertura que isto, definitivamente
func (f *SongFile) RetrieveSongText(s *Song) (*SongFile, error) {
	err := database.PsqlDB.
		Where("song_id = ? and type = ?", s.ID, text).
		First(f).Error

	if err != nil {
		return &SongFile{}, err
	}
	return f, err
}
