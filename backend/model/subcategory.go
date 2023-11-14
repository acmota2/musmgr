package model

import (
	"backend/db"

	"context"
)

type SubCategory struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  int64  `json:"category_id"`
}

func (s *SubCategory) Save() (SubCategory, error) {
	_, err := db.PsqlDB.Exec(
		context.Background(),
		`insert into subcategory(id, name, description, category_id) values($1, $2, $3, $4)`,
		s.ID,
		s.Name,
		s.Description,
		s.CategoryId,
	)
	if err != nil {
		return SubCategory{}, err
	}
	return *s, nil
}

func GetAllSongsFromSubcategory() (songs []Song, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		`select * from songs
		inner join song_subcategory on subcategory.id = song_subcategory.id`,
	)
	if err != nil {
		return []Song{}, err
	}

	for rows.Next() {
		var song Song
		err = rows.Scan(&song.ID, &song.Name, &song.Description, &song.Tonality)
		if err != nil {
			return []Song{}, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}
