package controllers

import (
	lib "main/library"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CalculateSegitiga(c *gin.Context) {

	alas, err := strconv.Atoi(c.Query("alas"))
	if err != nil {
		panic(err)
	}

	tinggi, err := strconv.Atoi(c.Query("tinggi"))
	if err != nil {
		panic(err)
	}

	calc := c.Query("hitung")

	var result gin.H

	segitiga := lib.Segitiga{Alas: alas, Tinggi: tinggi}
	if calc == "luas" {
		luas := segitiga.Luas()
		result = gin.H{
			"result": luas,
		}
	} else if calc == "keliling" {
		keliling := segitiga.Keliling()
		result = gin.H{
			"result": keliling,
		}

	}

	c.JSON(http.StatusOK, result)

}

func CalculatePersegi(c *gin.Context) {
	sisi, err := strconv.Atoi(c.Query("sisi"))
	if err != nil {
		panic(err)
	}

	calc := c.Query("hitung")

	var result gin.H

	persegi := lib.Persegi{Sisi: sisi}
	if calc == "luas" {
		luas := persegi.Luas()
		result = gin.H{
			"result": luas,
		}
	} else if calc == "keliling" {
		keliling := persegi.Keliling()
		result = gin.H{
			"result": keliling,
		}

	}

	c.JSON(http.StatusOK, result)

}

func CalculatePersegiPanjang(c *gin.Context) {
	panjang, err := strconv.Atoi(c.Query("panjang"))
	if err != nil {
		panic(err)
	}

	lebar, err := strconv.Atoi(c.Query("lebar"))
	if err != nil {
		panic(err)
	}

	calc := c.Query("hitung")

	var result gin.H

	persegiPanjang := lib.PersegiPanjang{Panjang: panjang, Lebar: lebar}
	if calc == "luas" {
		luas := persegiPanjang.Luas()
		result = gin.H{
			"result": luas,
		}
	} else if calc == "keliling" {
		keliling := persegiPanjang.Keliling()
		result = gin.H{
			"result": keliling,
		}

	}

	c.JSON(http.StatusOK, result)

}

func CalculateLingkaran(c *gin.Context) {
	jariJari, err := strconv.Atoi(c.Query("jariJari"))
	if err != nil {
		panic(err)
	}

	calc := c.Query("hitung")

	var result gin.H

	Lingkaran := lib.Lingkaran{JariJari: jariJari}
	if calc == "luas" {
		luas := Lingkaran.Luas()
		result = gin.H{
			"result": luas,
		}
	} else if calc == "keliling" {
		keliling := Lingkaran.Keliling()
		result = gin.H{
			"result": keliling,
		}

	}

	c.JSON(http.StatusOK, result)

}
