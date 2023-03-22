package main

import (
	"embed"
	"html/template"
	"net/http"

	"fast-api.io/mocks"
	"github.com/gin-gonic/gin"
)

//go:embed assets/* resources/*
var f embed.FS

func main() {
	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "resources/pages/*.html", "resources/base/*.html", "resources/sections/*.html"))
	router.SetHTMLTemplate(templ)

	router.StaticFS("/static", http.FS(f))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"slides": mocks.SLIDES_DATA,
			"works":  mocks.WORKS_DATA,
			"people": mocks.PEOPLE_DATA,
		})
	})

	router.GET("favicon.ico", func(c *gin.Context) {
		file, _ := f.ReadFile("assets/favicon.ico")
		c.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})

	router.Run(":8080")
}
