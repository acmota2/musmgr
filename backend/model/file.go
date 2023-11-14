package model

import (
	"context"

	"backend/db"
)

type FileType uint16

const (
	text  = iota
	score = iota
)

type SongFile struct {
	Path string   `json:"-"`
	Name string   `json:"name"`
	Open bool     `json:"-"`
	Type FileType `json:"type"`
}

func (f *SongFile) Save() (SongFile, error) {
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

func (f *SongFile) RetrieveSongText(songId int64) (_ SongFile, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		`select * from file
		where song_id = $1 and file_type = $2
		limit 1`,
		songId,
		text,
	)
	if err != nil {
		return SongFile{}, err
	}
	defer rows.Close()

	err = rows.Scan(&f.Name)
	if err != nil {
		return SongFile{}, err
	}
	return *f, err
}

func GetAllFilesFromSong(songId int64) (songFiles []SongFile, err error) {
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
