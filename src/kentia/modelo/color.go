package modelo

import (
	"kentia/log"

	"gopkg.in/mgo.v2/bson"
)

//Color es la estructura que definen los colores de la prenda
type Color struct {
	ID     bson.ObjectId `bson:"_id"`
	Tono   int
	Nombre string
}

const coleccionColor = "color"

//Registrar se encarga de registrar el color en la BD
func (c *Color) Registrar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionColor).Insert(c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//Modificar se encarga de modificar el color en la BD
func (c *Color) Modificar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionColor).UpdateId(c.ID, c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//ConsultarColores regresa un catálogo de colores
func ConsultarColores() (colores []Color) {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionColor).Find(bson.M{}).All(&colores)
	if err != nil {
		log.RegistrarError(err)
		return colores
	}
	return colores
}

//BuscarPorID busca en la BD un color que coincida con el ID dado
func (c *Color) BuscarPorID() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionColor).FindId(c.ID).One(c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}
