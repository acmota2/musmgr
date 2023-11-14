package model

import (
	"context"

	"backend/db"
)

type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetAllCategories() (categories []Category, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		"select 'id', 'name', 'description' from categories",
	)
	if err != nil {
		return []Category{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return []Category{}, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
