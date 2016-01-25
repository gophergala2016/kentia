package controlador

import (
	"fmt"
	"html/template"
	"kentia/genetico"
	"kentia/modelo"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2/bson"
)

//GenerarCombinacionGET maneja la ruta /generarCombinacion
func GenerarCombinacionGET(html *template.Template) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		usuarioID := GetSession(session.Get("UsuarioID"))
		if usuarioID != "0" {
			mapa := MapaInfo{}
			mapa.ObtenerDatosCombinacion(usuarioID.Hex())
			fmt.Println(mapa)
			html.ExecuteTemplate(c.Writer, "combinacion.html", mapa)
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
}

//GenerarMejorCombinacion se encarga de buscar cada una de las prendas por color y birllo para generar una combinacion.
func GenerarMejorCombinacion(usuarioID string) (prendas [][]modelo.Prenda) {
	u := modelo.Usuario{ID: bson.ObjectIdHex(usuarioID)}
	u.BuscarPorID()
	mejores := genetico.Genetico(u.ConsultarColoresPrendas())
	for _, mejor := range mejores {
		var combinacion []modelo.Prenda
		for i, color := range mejor.Genotipo {
			prenda := modelo.Prenda{}
			prenda.Brillo = color.Brillo
			prenda.Color.Tono = color.Tono
			switch i {
			case 0:
				prenda.TipoPrenda.Nombre = "Calzado"
			case 1:
				prenda.TipoPrenda.Nombre = "Pantalon/Falda"
			case 2:
				prenda.TipoPrenda.Nombre = "Playera"
			case 3:
				prenda.TipoPrenda.Nombre = "Chamarra"
			}
			prenda.BuscarPorBrilloTono(u.Prendas)
			combinacion = append(combinacion, prenda)
		}
		prendas = append(prendas, combinacion)
	}
	return prendas
}
