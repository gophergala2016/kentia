package controlador

import (
	"fmt"
	"html/template"
	"kentia/modelo"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegistroPrenda() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarioID := GetSession(sessions.Default(c).Get("UsuarioID"))
		if usuarioID != bson.ObjectId(0) {
			var p modelo.Prenda
			if c.Bind(&p) == nil {
				if p.Registar() {
					fmt.Println("BIEN", p)
				} else {
					fmt.Println("MAL", p)
				}
			} else {
				fmt.Println("Algo salió mal")
			}
			return
		}
		c.Redirect(302, "/")
		return
	}
}

func RegistroPrendaUsuario(html *template.Template) gin.HandlerFunc {
	return func(c *gin.Context) {
		mapa := MapaInfo{}
		mapa.ObtenerDatosRegistroPrenda()
		html.ExecuteTemplate(c.Writer, "registroPrenda.html", mapa)
	}
}
