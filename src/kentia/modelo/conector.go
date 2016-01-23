package modelo

import (
	"kentia/log"

	"gopkg.in/mgo.v2"
)

//Conector Go -> Mongo que permite obtener sesiones en una bd.
type conector struct {
	db *mgo.Database
}

// Realiza una conexión al servidor server y selecciona la BD de kentia
func (c *conector) IniciarSesion() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.RegistrarError(err)
	}
	session.SetSafe(&mgo.Safe{})
	c.db = session.DB("kentia")
}

// Finaliza la conexión al servidor
func (c *conector) CerrarSesion() {
	c.db.Logout()
}
