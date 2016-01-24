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
		mySession = bson.ObjectId(0)
	case bson.ObjectId:
		mySession = session.(bson.ObjectId)
	}
	return mySession
}

//Index handler para la ruta /
func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarioID := GetSession(sessions.Default(c).Get("UsuarioID"))
		if usuarioID != bson.ObjectId(0) {
			c.Redirect(http.StatusMovedPermanently, "/registroPrenda")
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}
