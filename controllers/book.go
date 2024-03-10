package controllers

import (
	"fmt"
	"main/database"
	"main/library"
	"main/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetAllBook(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book library.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	err = validator.New(validator.WithRequiredStructEnabled()).Struct(book)
	if err != nil {
		fmt.Println(err)
	}

	book.Thickness = library.GetBookThickness(book.Total_page)

	err = validator.New(validator.WithRequiredStructEnabled()).Struct(book)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"result": err.Error(),
		})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Book",
	})
}

func UpdateBook(c *gin.Context) {
	var book library.Book
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	err = c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	err = validator.New(validator.WithRequiredStructEnabled()).Struct(book)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"result": err.Error(),
		})
		panic(err)
	}

	book.Id = id
	book.Updated_at = library.GetCurrentTimestamp()
	book.Thickness = library.GetBookThickness(book.Total_page)

	err = repository.UpdateBook(database.DbConnection, book)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Book",
	})
}

func DeleteBook(c *gin.Context) {
	var book library.Book
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	book.Id = id
	err = repository.DeleteBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Data",
	})

}

func GetAllBookByCategory(c *gin.Context) {
	var (
		result   gin.H
		category library.Category
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	category.Id = id
	books, err := repository.GetAllBookByCategory(database.DbConnection, category)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}
