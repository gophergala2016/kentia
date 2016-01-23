package main

import (
	"html/template"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var servidor *gin.Engine

func init() {
	servidor = gin.Default()
	gin.SetMode(gin.DebugMode)
}

func main() {
	html := template.Must(template.ParseFiles("something.html"))
	servidor.SetHTMLTemplate(html)

	servidor.Use(static.Serve("/", static.LocalFile("./other", false)))
	servidor.Run(":3000")
}
