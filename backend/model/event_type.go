package model

import (
	"backend/db"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type EventType struct {
	Name string `json:"name"`
}

func GetAllEventTypes() (eventTypes []EventType, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		"select * from event_type",
	)
	if err != nil {
		return []EventType{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var eventType EventType
		err = rows.Scan(&eventType.Name)
		if err != nil {
			return []EventType{}, err
		}
		eventTypes = append(eventTypes, eventType)
	}

	if len(eventTypes) == 0 {
		return eventTypes, pgx.ErrNoRows
	}
	return eventTypes, err
}

func GetAllEventsFromEventType(eventName string) (events []Event, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		`select id, description, date from event
		where event_type_name = $1`,
		eventName,
	)
	if err != nil {
		return []Event{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		var date pgtype.Date
		err = rows.Scan(&event.ID, &event.Description, &date)
		if err != nil {
			return []Event{}, err
		}
		event.Date = DayOfYear(date.Time)
		events = append(events, event)
	}
	return events, nil
}
