package model

import (
	"backend/db"
	"context"
)

type SongEvent struct {
	SongID  int64 `json:"song_id"`
	EventID int64 `json:"event_id"`
}

func (se *SongEvent) Save() (SongEvent, error) {
	err := db.PsqlDB.QueryRow(
		context.Background(),
		`insert into song_event(song_id, event_id)
        values ($1, $2) returning event_id`,
		se.SongID,
		se.EventID,
	).Scan(&se.EventID)
	if err != nil {
		return SongEvent{}, err
	}
	return *se, nil
}

// future: GetAllEventsFromSong
