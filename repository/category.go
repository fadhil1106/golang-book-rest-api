package repository

import (
	"database/sql"
	lib "main/library"
)

func GetAllCategory(db *sql.DB) (results []lib.Category, err error) {
	sql := "SELECT * from category"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var category = lib.Category{}

		err = rows.Scan(&category.Id, &category.Name, &category.Created_at, &category.Updated_at)
		if err != nil {
			panic(err)
		}
		results = append(results, category)
	}

	return
}

func InsertCategory(db *sql.DB, category lib.Category) (err error) {
	sql := "INSERT INTO category (name) VALUES ($1)"

	errs := db.QueryRow(sql, category.Name)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category lib.Category) (err error) {
	sql := "UPDATE category SET name = $1, updated_at = $2 WHERE id = $3"
	errs := db.QueryRow(sql, category.Name, category.Updated_at, category.Id)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category lib.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"

	errs := db.QueryRow(sql, category.Id)

	return errs.Err()
}
