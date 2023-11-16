package model

import (
	"context"

	"backend/db"
)

type Song struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Tonality       string  `json:"tonality"`
	SubCategoryIds []int64 `json:"subcategories"`
}

func (s *Song) Save() (Song, error) {
	err := db.PsqlDB.QueryRow(
		context.Background(),
		`insert into song(name, tonality)
		values ($1,$2)
		returning id`,
		s.Name,
		s.Tonality,
	).Scan(&s.ID)

	if err != nil {
		return Song{}, err
	}

	for _, id := range s.SubCategoryIds {
		_, err = db.PsqlDB.Exec(
			context.Background(),
			`insert into song_subcategory(song_id, subcategory_id) values ($1, $2)`,
			s.ID,
			id,
		)
		if err != nil {
			return Song{}, err
		}
	}
	return *s, nil
}

func GetAllSongs() (songs []Song, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		`select * from song`,
	)
	if err != nil {
		return []Song{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var song Song
		err = rows.Scan(&song.ID, &song.Name, &song.Tonality)
		if err != nil {
			return []Song{}, nil
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func GetSongSubcategories(songId int64) (subCats []SubCategory, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		`select id, name, category_id from subcategory
		inner join song_subcategory on song_subcategory.subcategory_id = $1`,
		songId,
	)
	if err != nil {
		return []SubCategory{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var subCat SubCategory
		err = rows.Scan(&subCat.ID, &subCat.Name, &subCat.CategoryId)
		if err != nil {
			return []SubCategory{}, nil
		}
		subCats = append(subCats, subCat)
	}
	return subCats, nil
}
