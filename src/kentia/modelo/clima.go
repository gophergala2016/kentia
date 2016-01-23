package modelo

import (
	"kentia/log"

	"gopkg.in/mgo.v2/bson"
)

//Clima es la estructura que define los climas para los que se usara la prenda
type Clima struct {
	ID     bson.ObjectId `bson:"_id" form:"id"`
	Nombre string
}

const coleccionClima = "clima"

//Registrar se encarga de registrar el clima en la BD
func (c *Clima) Registrar() bool {
	var conn conector
	conn.IniciarSesion()
	defer conn.CerrarSesion()
	err := conn.db.C(coleccionClima).Insert(c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}
