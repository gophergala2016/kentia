package controlador

import (
	"html/template"
	"kentia/modelo"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*Login EJEMPLO DE login. Funcion que manejara la ruta POST /login*/
func Login(html *template.Template) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		usuario := modelo.Usuario{}
		c.Bind(&usuario)

		isOk := usuario.IniciarSesion()
		if !isOk {
			session.Set("UsuarioID", bson.ObjectId(0))
		}
		session.Set("UsuarioID", usuario.ID)
		session.Save()
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}
