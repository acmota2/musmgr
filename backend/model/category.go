package model

import (
	"context"
	"fmt"

	"backend/db"
)

type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetAllCategories() (categories []Category, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		`select * from category`,
	)
	if err != nil {
		return []Category{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		fmt.Print("estoura aí, oh mano\n")
		err := rows.Scan(&category.Name, &category.Description)
		if err != nil {
			return []Category{}, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func GetAllSubCategoriesFromCategory(categoryName string) (subCats []SubCategory, err error) {
	rows, err := db.PsqlDB.Query(
		context.Background(),
		`select * from subcategory
		where category_name = $1`,
		categoryName,
	)
	if err != nil {
		return []SubCategory{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var subCat SubCategory
		err := rows.Scan(&subCat.ID, &subCat.Name, &subCat.CategoryName)
		if err != nil {
			return []SubCategory{}, err
		}
		subCats = append(subCats, subCat)
	}
	return subCats, nil
}
