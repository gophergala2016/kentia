package modelo

import (
	"kentia/log"

	"gopkg.in/mgo.v2/bson"
)

//Clima es la estructura que define los climas para los que se usara la prenda
type Clima struct {
	ID     bson.ObjectId `bson:"_id"`
	Nombre string
}

const coleccionClima = "clima"

//Registrar se encarga de registrar el clima en la BD
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

//Modificar
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

//ConsultarClima regresa un catálogo de colores
func ConsultarClima() (climas []Clima) {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionClima).Find(bson.M{}).All(&climas)
	if err != nil {
		log.RegistrarError(err)
	}
	return climas
}

//BuscarPorID busca en la BD un cllima que coincida con el ID dado
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
