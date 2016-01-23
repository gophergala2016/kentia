package controlador

import (
	"kentia/modelo"

	"github.com/gin-gonic/gin"
)

//RegistroUsuario procesa los datos recibidos del formulario
func RegistroUsuario() gin.HandlerFunc {
	return func(c *gin.Context) {
		var u modelo.Usuario
		if c.Bind(&u) == nil {
			if u.Registrar() {
				//Correcto
			} else {
				//Algo salio mal
			}
		}
	}
}
