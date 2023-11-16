package model

import (
	"backend/db"

	"context"
)

type SubCategory struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	CategoryId int64  `json:"category_id"`
}

func (s *SubCategory) Save() (SubCategory, error) {
	_, err := db.PsqlDB.Exec(
		context.Background(),
		`insert into subcategory(name, category_id) values($2, $3)`,
		s.Name,
		s.CategoryId,
	)
	if err != nil {
		return SubCategory{}, err
	}
	return *s, nil
}

func GetAllSongsFromSubcategory(subcategoryId int64) (songs []Song, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		`select id, name, tonality from song
		inner join song_subcategory on song_subcategory.subcategory_id = $1`,
		subcategoryId,
	)
	if err != nil {
		return []Song{}, err
	}

	for rows.Next() {
		var song Song
		err = rows.Scan(&song.ID, &song.Name, &song.Tonality)
		if err != nil {
			return []Song{}, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func GetAllSubCategories() (subCats []SubCategory, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		`select * from subcategory`,
	)
	if err != nil {
		return []SubCategory{}, nil
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
