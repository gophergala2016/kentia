package modelo

import (
	"kentia/log"

	"gopkg.in/mgo.v2/bson"
)

//tipo de clima para el que se usa esta prenda
type Ocasion struct {
	ID     bson.ObjectId `bson:"_id"`
	Nombre string
}

const coleccionOcasion = "ocacion"

func (c *Ocasion) Registar() bool {
	conn := conectar()
	defer conn.desconectar()
err:
	dao.db.C(coleccionOcacion).Insert(c)
	if err != nil {
		log.RegistarError(err)
		return false
	}
	return true
}

//Modificar
func (c *Ocasion) Modificar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionOcasion).UpdateId(c.ID, c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//ConsultarOcasion regresa un catálogo de colores
func ConsultarOcasion() (ocasiones []Ocasion) {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionOcasion).Find(bson.M{}).All(&ocasiones)
	if err != nil {
		log.RegistrarError(err)
		return ocasiones
	}
	return ocasiones
}

//BuscarPorID busca en la BD un cllima que coincida con el ID dado
func (c *Ocasion) BuscarPorID() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionOcasion).FindId(c.ID).One(c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}
