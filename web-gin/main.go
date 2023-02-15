package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"title": "Main Website",
			"date":  time.Date(2023, 2, 14, 15, 13, 14, 0, time.Local),
		})
	})
	r.GET("/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
			"posts": []string{"Go", "PHP", "Java", "JavaScript"},
		})
	})
	r.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
			"users": []string{},
		})
	})
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})
	r.POST("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
	r.PUT("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	r.DELETE("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})
	_ = r.Run()
}
