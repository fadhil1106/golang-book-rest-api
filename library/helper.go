package library

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(auth gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete {
			auth(c)
		}
	}
}

func GetCurrentTimestamp() string {
	now := time.Now()
	time := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	return time
}

func GetBookThickness(totalPage int) string {
	var thickness string
	if totalPage <= 100 {
		thickness = "tipis"
	} else if totalPage > 100 && totalPage <= 200 {
		thickness = "sedang"
	} else {
		thickness = "tebal"
	}

	return thickness
}

func IsValidUrl(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return false
	}
	return true
}

func IsValidReleaseYear(year int) bool {
	if year >= 1980 && year <= 2021 {
		return true
	} else {
		return false
	}
}
