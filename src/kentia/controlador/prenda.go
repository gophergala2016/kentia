package controlador

import (
	"fmt"
	"html/template"
	"kentia/modelo"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegistroPrendaPOST() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarioID := GetSession(sessions.Default(c).Get("UsuarioID"))
		if usuarioID != bson.ObjectId(0) {
			var p modelo.Prenda
			if c.Bind(&p) == nil {
				u := modelo.Usuario{ID: usuarioID}
				if u.BuscarPorID() {
					p.ID = bson.NewObjectId()
					p.Clima.BuscarPorID()
					p.Color.BuscarPorTono()
					p.Ocasion.BuscarPorID()
					p.TipoPrenda.BuscarPorID()
					u.Prendas = append(u.Prendas, p)
					if u.Modificar() {
						//BIEN
						fmt.Println(u)
					} else {
						fmt.Println("ALGO MAL", u)
					}
				} else {
					//No se encontró el usuario D:
					fmt.Println(u)
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

func RegistroPrendaGET(html *template.Template) gin.HandlerFunc {
	return func(c *gin.Context) {
		mapa := MapaInfo{}
		mapa.ObtenerDatosRegistroPrenda()
		html.ExecuteTemplate(c.Writer, "registroPrenda.html", mapa)
	}
}
