package main

import (
	"database/sql"
	"fmt"
	"main/controllers"
	"main/database"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// ENV Configuration
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success read file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOSTNAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println(psqlInfo)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()

	auth := gin.BasicAuth(gin.Accounts{
		"admin":  "password",
		"editor": "secret",
	})

	// ROUTER GIN BANGUN DATAR
	router.GET("/bangung-datar/segitiga-sama-sisi", controllers.CalculateSegitiga)
	router.GET("/bangung-datar/persegi", controllers.CalculatePersegi)
	router.GET("/bangung-datar/persegi-panjang", controllers.CalculatePersegiPanjang)
	router.GET("/bangung-datar/lingkaran", controllers.CalculateLingkaran)

	// ROUTER GIN CATEGORY
	router.GET("/categories", controllers.GetAllCategory)
	router.POST("/categories", AuthMiddleware(auth), controllers.InsertCategory)
	router.PUT("/categories/:id", AuthMiddleware(auth), controllers.UpdateCategory)
	router.DELETE("/categories/:id", AuthMiddleware(auth), controllers.DeleteCategory)
	router.GET("/categories/:id/books", controllers.GetAllBookByCategory)

	// ROUTER GIN BOOK
	router.GET("/books", controllers.GetAllBook)
	router.POST("/books", AuthMiddleware(auth), controllers.InsertBook)
	router.PUT("/books/:id", AuthMiddleware(auth), controllers.UpdateBook)
	router.DELETE("/books/:id", AuthMiddleware(auth), controllers.DeleteBook)

	router.Run("localhost:8888")
}

func AuthMiddleware(auth gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete {
			auth(c)
		}
	}
}
