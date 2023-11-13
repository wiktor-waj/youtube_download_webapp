package main

import (
	"html/template"
	"net/http"

	"github.com/wiktor-waj/youtube_download_webapp/page"

	"github.com/gin-gonic/gin"
)

func viewHandler(c *gin.Context) {
	p, err := page.LoadPage("index.html")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Writer.Write(p.Body)
}

func main() {
	html := template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))
	router := gin.Default()
	router.SetHTMLTemplate(html)
	router.SetTrustedProxies(nil)
	router.Static("/css", "data/css")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", viewHandler)

	router.Run(":8080")
}
