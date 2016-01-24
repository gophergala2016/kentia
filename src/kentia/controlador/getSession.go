package controlador

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

//GetSession obtiene el valor bson.ObjectId de la sesion actual*/
func GetSession(session interface{}) bson.ObjectId {
	var mySession bson.ObjectId
	switch session.(type) {
	default:
		mySession = "0"
	case string:
		mySession = bson.ObjectIdHex(session.(string))
	}
	return mySession
}

//Index handler para la ruta /
func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		usuarioID := GetSession(session.Get("UsuarioID"))
		if usuarioID != "0" {
			c.Redirect(http.StatusTemporaryRedirect, "/registroPrenda")
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}
