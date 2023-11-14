package model

import (
	"context"

	"backend/db"
)

type Song struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Tonality       string  `json:"tonality"`
	SubCategoryIds []int64 `json:"subcategories"`
}

func (s *Song) Save() (Song, error) {
	_, err := db.PsqlDB.Exec(
		context.Background(),
		`insert into song(id, name, description, tonality) values ($1,$2,$3,$4)`,
		s.ID,
		s.Name,
		s.Description,
		s.Tonality,
	)
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
