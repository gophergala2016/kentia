package main

import (
	"html/template"
	"kentia/controlador"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	html     *template.Template
	servidor *gin.Engine
)

func init() {
	gin.SetMode(gin.DebugMode)
	servidor = gin.Default()
	cargarTemplates()
	servidor.Use(static.Serve("/", static.LocalFile("./public", false)))
	servidor.StaticFile("/", "./public/index.html")
	servidor.NoRoute(func(c *gin.Context) {
		html.ExecuteTemplate(c.Writer, "404.html", nil)
	})
}

func cargarTemplates() {
	html = template.Must(template.ParseFiles("public/404.html"))
	servidor.SetHTMLTemplate(html)
}

func main() {
	servidor.StaticFile("/RegistroUsuario", "./public/registro.html")

	servidor.POST("/registroUsuario", controlador.RegistroUsuario())
	servidor.Run(":3000")
}
