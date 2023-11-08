package model

import (
	database "backend/db"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name string
	Date *time.Time
}

func (e *Event) Save() (*Event, error) {
	err := database.DB.Create(e).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}
