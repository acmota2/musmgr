package model

import (
	"context"
	"fmt"

	"backend/db"
)

type FileType uint16

const (
	text  = iota
	score = iota
)

type SongFile struct {
	Path   string   `json:"-"`
	Name   string   `json:"name"`
	Open   bool     `json:"-"`
	SongId int64    `json:"song_id"`
	Type   FileType `json:"type"`
}

func (f *SongFile) Save() (SongFile, error) {
	path := fmt.Sprintf("./song_files/%s.json", f.Name)
	f.Path = path
	_, err := db.PsqlDB.Exec(
		context.Background(),
		`insert into files(path, name, open, type) values ($1, $2, $3, $4)`,
		f.Path,
		f.Name,
		false,
		f.Type,
	)
	if err != nil {
		return SongFile{}, err
	}

	return *f, nil
}

func GetTextFileFromSong(songId int64) (SongFile, error) {
	row := db.PsqlDB.QueryRow(
		context.Background(),
		`select * from file
		where song_id = $1 and type = $2`,
		songId,
		text,
	)
	var songFile SongFile
	if err := row.Scan(&songFile.Path, &songFile.Name, &songFile.Open, &songFile.Type); err != nil {
		return SongFile{}, err
	}
	return songFile, nil
}

func GetAllFilesFromSong(songId int64) (songFiles []SongFile, _ error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		`select * from file
		where song_id = $1`,
		songId,
	)
	if err != nil {
		return []SongFile{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var songFile SongFile
		err = rows.Scan(&songFile.Path, &songFile.Name, &songFile.Open, &songFile.Type)
		if err != nil {
			return []SongFile{}, nil
		}
		songFiles = append(songFiles, songFile)
	}
	return songFiles, nil
}
