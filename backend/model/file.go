package model

import (
	"context"
	"fmt"
	"os"

	"backend/db"
)

type FileType uint16

const (
	text  = iota
	score = iota
)

type SongFile struct {
	Path     string   `json:"-"`
	Name     string   `json:"name"`
	Open     bool     `json:"open"`
	SongId   int64    `json:"song_id"`
	Type     FileType `json:"type"`
	TextFile string   `json:"text_file"`
}

func (f *SongFile) Save() (SongFile, error) {
	path := fmt.Sprintf("./song_files/%s.json", f.Name)
	f.Path = path

	err := os.WriteFile(f.Path, []byte(f.TextFile), 0666)
	if err != nil {
		return SongFile{}, err
	}

	_, err = db.PsqlDB.Exec(
		context.Background(),
		`insert into files(path, name, open, type, song_id) 
        values ($1, $2, $3, $4, $5)`,
		f.Path,
		f.Name,
		false,
		f.Type,
		f.SongId,
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

	content, err := os.ReadFile(songFile.Path)
	if err != nil {
		return SongFile{}, err
	}

	songFile.TextFile = string(content)
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
