package repository

import (
	"database/sql"
	lib "main/library"
)

func GetAllBook(db *sql.DB) (results []lib.Book, err error) {
	sql := `
	SELECT books.*, category.*
	FROM books
	INNER JOIN category
	ON books.category_id = category."id"
	`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var book = lib.Book{}

		err = rows.Scan(&book.Id, &book.Title, &book.Description, &book.Image_url,
			&book.Release_year, &book.Price, &book.Total_page, &book.Thickness, &book.Created_at, &book.Updated_at, &book.Category_id,
			&book.Category.Id, &book.Category.Name, &book.Category.Created_at, &book.Category.Updated_at)
		if err != nil {
			panic(err)
		}
		results = append(results, book)
	}

	return
}

func InsertBook(db *sql.DB, book lib.Book) (err error) {
	sql := `INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	errs := db.QueryRow(sql, book.Title, book.Description, book.Image_url, book.Release_year, book.Price, book.Total_page, book.Thickness, book.Category_id)

	return errs.Err()
}

func UpdateBook(db *sql.DB, book lib.Book) (err error) {
	sql := `UPDATE books SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, 
				category_id=$8, updated_at = $9 WHERE id = $10`
	errs := db.QueryRow(sql, book.Title, book.Description, book.Image_url, book.Release_year, book.Price, book.Total_page, book.Thickness, book.Category_id, book.Updated_at, book.Id)

	return errs.Err()
}

func DeleteBook(db *sql.DB, book lib.Book) (err error) {
	sql := "DELETE FROM books WHERE id = $1"

	errs := db.QueryRow(sql, book.Id)

	return errs.Err()
}

func GetAllBookByCategory(db *sql.DB, category lib.Category) (results []lib.Book, err error) {
	sql := `
	SELECT
		books.*, 
		category.*
	FROM
		books
		INNER JOIN
		category
		ON 
			books.category_id = category."id"
	WHERE
		books.category_id = $1
	`

	rows, err := db.Query(sql, category.Id)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var book = lib.Book{}

		err = rows.Scan(&book.Id, &book.Title, &book.Description, &book.Image_url,
			&book.Release_year, &book.Price, &book.Total_page, &book.Thickness, &book.Created_at, &book.Updated_at, &book.Category_id,
			&book.Category.Id, &book.Category.Name, &book.Category.Created_at, &book.Category.Updated_at)
		if err != nil {
			panic(err)
		}
		results = append(results, book)
	}

	return
}
