package model

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"backend/db"
	"backend/songs"
)

type FileType uint16

const (
	text  = iota
	score = iota
)

type SongFile struct {
	Path     string          `json:"-"`
	Name     string          `json:"name"`
	Open     bool            `json:"-"`
	Type     FileType        `json:"type"`
	TextFile *songs.SongText `json:"song_text"`
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

	file, err := os.Open(path)
	if err != nil {
		return SongFile{}, err
	}

	fileJSON, err := json.MarshalIndent(f.TextFile, "", "    ")
	if err != nil {
		return SongFile{}, err
	}

	_, err = file.Write(fileJSON)
	if err != nil {
		return SongFile{}, err
	}

	return *f, nil
}

func (f *SongFile) RetrieveSongText(songId int64) (_ songs.SongText, err error) {
	row := db.PsqlDB.QueryRow(
		context.Background(),
		`select * from file
		where song_id = $1 and file_type = $2`,
		songId,
		text,
	)
	if err != nil {
		return songs.SongText{}, err
	}

	err = row.Scan(&f.Name)
	if err != nil {
		return songs.SongText{}, err
	}
	return *f.TextFile, err
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
