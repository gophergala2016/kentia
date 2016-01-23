package main

import (
	"html/template"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	servidor := gin.Default()

	html := template.Must(template.ParseFiles("something.html"))
	servidor.SetHTMLTemplate(html)

	servidor.Use(static.Serve("/", static.LocalFile("./other", false)))
	servidor.Run(":3000")
}
