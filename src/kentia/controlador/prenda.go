package controlador

import (
	"html/template"
	"kentia/modelo"

	"github.com/gin-gonic/gin"
)

func RegistroPrendaPOST() gin.HandlerFunc {
	return func(c *gin.Context) {
		prenda := modelo.Prenda{}
		c.Bind(prenda)

	}
}

func RegistroPrendaGET(html *template.Template) gin.HandlerFunc {
	return func(c *gin.Context) {
		mapa := MapaInfo{}
		mapa.ObtenerDatosRegistroPrenda()
		html.ExecuteTemplate(c.Writer, "registroPrenda.html", mapa)
	}
}
