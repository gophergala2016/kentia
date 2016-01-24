package modelo

import (
	"kentia/log"

	"gopkg.in/mgo.v2/bson"
)

//Clima es la estructura que define los climas para los que se usará la prenda.
type Clima struct {
	ID     bson.ObjectId `bson:"_id"`
	Nombre string
}

const coleccionClima = "clima"

//Registrar se encarga de registrar el clima en la BD.
func (c *Clima) Registrar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionClima).Insert(c)
	if err != nil {
		log.RegistrarError(err)

		return false
	}
	return true
}

//Modificar actualiza los datos del clima.
func (c *Clima) Modificar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionClima).UpdateId(c.ID, c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//ConsultarClimas regresa un catálogo de colores
func ConsultarClimas() (climas []Clima) {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionClima).Find(bson.M{}).All(&climas)
	if err != nil {
		log.RegistrarError(err)
	}
	return climas
}

//BuscarPorID busca en la BD un clima que coincida con el ID dado.
func (c *Clima) BuscarPorID() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionClima).FindId(c.ID).One(c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}
