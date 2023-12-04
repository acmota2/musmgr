package model

import (
	"backend/db"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Event struct {
	ID            int64     `json:"id"`
	Date          DayOfYear `json:"date"`
	Description   string    `json:"description"`
	EventTypeName string    `json:"event_type_name"`
}

func (e *Event) Save() (Event, error) {
	err := db.PsqlDB.QueryRow(
		context.Background(),
		`insert into event(id, date, description, event_type_name)
		values (default, $1, $2, $3) returning id`,
		time.Time(e.Date),
		e.Description,
		e.EventTypeName,
	).Scan(&e.ID)
	if err != nil {
		return Event{}, err
	}
	return *e, nil
}

func GetAllEvents() (events []Event, _ error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		"select * from event",
	)
	if err != nil {
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		var date time.Time
		err := rows.Scan(&event.ID, &date, &event.Description, &event.EventTypeName)
		event.Date = DayOfYear(date)
		if err != nil {
			return []Event{}, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventSongs(eventId int64) (songs []Song, _ error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		`select id, name, tonality from song
		inner join song_event on song.id = song_event.song_id
		where song.id = $1`,
		eventId,
	)
	if err != nil {
		return []Song{}, err
	}

	for rows.Next() {
		var song Song
		err := rows.Scan(&song.ID, &song.Name, &song.Tonality)
		if err != nil {
			return []Song{}, err
		}
		songs = append(songs, song)
	}

	if len(songs) == 0 {
		return []Song{}, pgx.ErrNoRows
	}
	return songs, nil
}
