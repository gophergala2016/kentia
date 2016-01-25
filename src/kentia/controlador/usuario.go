package controlador

import (
	"fmt"
	"html/template"
	"kentia/modelo"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//Login es la función que se encarga de comprobar las credenciales de los usuarios.
func Login(html *template.Template) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		usuario := modelo.Usuario{}
		c.Bind(&usuario)

		isOk := usuario.IniciarSesion()
		if !isOk {
			session.Set("UsuarioID", "0")
		}
		session.Set("UsuarioID", usuario.ID.Hex())
		session.Save()
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
}

//RegistroUsuario procesa los datos recibidos del formulario.
func RegistroUsuario() gin.HandlerFunc {
	return func(c *gin.Context) {
		var u modelo.Usuario
		if c.Bind(&u) == nil {
			if u.Registrar() {
				//Correcto
				session := sessions.Default(c)
				session.Set("UsuarioID", u.ID.Hex())
				session.Save()
				c.Redirect(http.StatusSeeOther, "/")
				return
			}
			//Algo salio mal
			c.String(http.StatusInternalServerError, "No se pudo registrar")
			fmt.Println("No registrado ", u)
		} else {
			c.Redirect(http.StatusSeeOther, "/")
			fmt.Println("Datos incorrectos")
		}
	}
}
