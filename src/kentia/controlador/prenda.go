package controlador

import (
	"fmt"
	"html/template"
	"kentia/modelo"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func convertirID(s string) string {
	s = s[strings.Index(s, "\"")+1 : strings.LastIndex(s, "\"")]
	return s
}

func RegistroPrendaPOST() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarioID := GetSession(sessions.Default(c).Get("UsuarioID"))
		if usuarioID != "0" {
			var p modelo.Prenda
			if c.Bind(&p) == nil {
				u := modelo.Usuario{ID: usuarioID}
				if u.BuscarPorID() {
					p.ID = bson.NewObjectId()
					p.Color.BuscarPorTono()

					p.Clima.ID = bson.ObjectIdHex(convertirID(c.PostForm("clima")))
					p.Clima.BuscarPorID()

					p.TipoPrenda.ID = bson.ObjectIdHex(convertirID(c.PostForm("tipoPrenda")))
					p.TipoPrenda.BuscarPorID()

					p.Ocasion.ID = bson.ObjectIdHex(convertirID(c.PostForm("ocasion")))
					p.Ocasion.BuscarPorID()
					fmt.Println(p)

					fmt.Println(p)
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
