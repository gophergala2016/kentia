package controlador

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"kentia/modelo"
	"net/http"
	"os"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func convertirID(s string) string {
	s = s[strings.Index(s, "\"")+1 : strings.LastIndex(s, "\"")]
	return s
}

func guadarImagen(c *gin.Context, p *modelo.Prenda) {
	file, _, err := c.Request.FormFile("foto")
	if err != nil {
		c.String(http.StatusSeeOther, "Sin imagen")
		return
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)

	ruta := "/img/foto" + p.ID.Hex() + ".png"
	p.Foto = ruta

	out, err := os.Create("public" + p.Foto)
	if err != nil {
		c.String(http.StatusTemporaryRedirect, err.Error())
		return
	}

	_, err = out.Write(data)
	if err != nil {
		c.String(http.StatusTemporaryRedirect, err.Error())
		return
	}
	defer out.Close()
}

//RegistroPrendaPOST recibe el formulario y se encarga de registrarlo en la BD.
func RegistroPrendaPOST() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarioID := GetSession(sessions.Default(c).Get("UsuarioID"))
		if usuarioID != "0" {
			var p modelo.Prenda
			err := c.Bind(&p)
			if err == nil {
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

					guadarImagen(c, &p)

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
				fmt.Println("Algo salió mal", err)
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

func MuestraPrendasGET(html *template.Template) gin.HandlerFunc {
	return func(c *gin.Context) {
		mapa := MapaInfo{}
		usuarioID := GetSession(sessions.Default(c).Get("UsuarioID"))
		fmt.Println(usuarioID)
		mapa.ObtenerDatosPrendas(usuarioID)
		html.ExecuteTemplate(c.Writer, "principal.html", mapa)
	}
}
