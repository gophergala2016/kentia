package modelo

import (
	"fmt"
	"kentia/log"

	"gopkg.in/mgo.v2/bson"
)

//Ocasion estructura para conocer la ocasion en que se usará la prenda.
type Ocasion struct {
	ID     bson.ObjectId `bson:"_id"`
	Nombre string
}

const coleccionOcasion = "ocasion"

//Registrar un nuevo tipo de ocacion en la bd.
func (c *Ocasion) Registrar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionOcasion).Insert(c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//Modificar actualizara la ocasión.
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

//ConsultarOcasiones regresa un catálogo de ocasiones.
func ConsultarOcasiones() (ocasiones []Ocasion) {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionOcasion).Find(bson.M{}).All(&ocasiones)
	if err != nil {
		log.RegistrarError(err)
	}
	fmt.Println(ocasiones)

	return ocasiones
}

//BuscarPorID busca en la BD un ocasion que coincida con el ID dado.
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
