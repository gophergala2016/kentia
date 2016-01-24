package modelo

import (
	"kentia/log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Combinacion es la estructura que definen los colores de la prenda
type Combinacion struct {
	ID       bson.ObjectId `bson:"_id"`
	Prendas  []Prenda
	FechaUso []time.Time
	Favorito bool
}

const coleccionCombinacion = "combinacion"

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

//ConsultarPorID para consultar una combinacion por ID
func (c *Combinacion) ConsultarPorID() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionCombinacion).FindId(c.ID).One(c)
	if err != nil {
		log.RegistrarError(err)
		return true
	}
	return false
}

//Modificar se encarga de modificar la combinacion en la BD
func (c *Combinacion) Modificar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionCombinacion).UpdateId(c.ID, c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}
