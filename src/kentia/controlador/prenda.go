package controlador

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func RegistroPrenda() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func RegistroPrendaUsuario(html *template.Template) gin.HandlerFunc {
	return func(c *gin.Context) {
		mapa := MapaInfo{}
		mapa.ObtenerDatosRegistroPrenda()
		html.ExecuteTemplate(c.Writer, "registroPrenda.html", mapa)
	}
}
