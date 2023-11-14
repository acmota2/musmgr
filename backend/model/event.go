package model

import (
	"context"
	"time"

	"backend/db"
)

type Event struct {
	ID   int64      `json:"id"`
	Name string     `json:"name"`
	Date *time.Time `json:"date"`
}

func (e *Event) Save() (Event, error) {
	_, err := db.PsqlDB.Exec(
		context.Background(),
		"insert event('name', 'date') values ($1, $2);",
		e.Name,
		e.Date,
	)
	if err != nil {
		return Event{}, err
	}
	return *e, nil
}

func GetAllEvents() (events []Event, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		"select id, name, date from event",
	)
	if err != nil {
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Date)
		if err != nil {
			return []Event{}, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventSongs(eventId int64) (songs []Song, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		`select * from song
		inner join song_event on song.id = song_event.id
		where song.id = $1`,
		eventId,
	)
	if err != nil {
		return []Song{}, err
	}

	for rows.Next() {
		var song Song
		err := rows.Scan(&song.ID, &song.Name, &song.Description, &song.Tonality)
		if err != nil {
			return []Song{}, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}
