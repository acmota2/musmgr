package model

import (
	database "backend/db"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name  string
	Date  *time.Time
	Songs []*Song `json:"songs"`
}

func (e *Event) Save() (*Event, error) {
	err := database.PsqlDB.Create(e).Error
	if err != nil {
		return &Event{}, err
	}
	return e, nil
}

func GetAllEvents() (*[]Event, error) {
	var events *[]Event
	err := database.PsqlDB.Find(&events).Error

	if err != nil {
		return &[]Event{}, err
	}
	return events, nil
}

func GetEventSongs(eventId uint64) ([]*Song, error) {
	var event *Event
	err := database.PsqlDB.Model(&Event{}).Preload("songs").First(&event, eventId).Error

	if err != nil {
		return []*Song{}, err
	}
	return event.Songs, nil
}
