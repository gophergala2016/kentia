package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var servidor *gin.Engine

func init() {
	gin.SetMode(gin.DebugMode)
	servidor = gin.Default()
	servidor.Use(static.Serve("/", static.LocalFile("./public", false)))
	servidor.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusOK, "index.html")
	})
}

func main() {
	html := template.Must(template.ParseFiles("something.html"))
	servidor.SetHTMLTemplate(html)

	servidor.Run(":3000")
}
