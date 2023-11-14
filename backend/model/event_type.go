package model

import (
	"context"

	"backend/db"
)

type EventType struct {
	Name string `json:"name"`
}

func GetEventTypes() (eventTypes []EventType, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		"select 'name' from event_type",
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
	return eventTypes, err
}
