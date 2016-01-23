package controlador

import (
	"fmt"
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
				fmt.Println("Registrado ", u)
			} else {
				//Algo salio mal
				fmt.Println("No registrado ", u)
			}
		} else {
			fmt.Println("Datos incorrectos")
		}
	}
}
