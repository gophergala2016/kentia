package modelo

import (
	"kentia/log"

	"gopkg.in/mgo.v2/bson"
)

//Combinacion es la estructura que definen los colores de la prenda
type Combinacion struct {
	ID     bson.ObjectId `bson:"_id" form:"id"`
	Nombre string
}

const coleccionCombinacion = "color"

//Registrar se encarga de registrar el color en la BD
func (c *Combinacion) Registrar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionCombinacion).Insert(c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}
