package modelo

import (
	"kentia/log"

	"gopkg.in/mgo.v2/bson"
)

//Prenda .
type Prenda struct {
	ID         bson.ObjectId `bson:"_id"`
	Brillo     int           `form:"luminucidad"`
	Foto       string
	Color      Color
	Clima      Clima
	Tipoprenda TipoPrenda
	Ocasion    Ocasion
}

const coleccionPrenda = "prenda"

//Registar funcion para cargar clima
func (c *Prenda) Registar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionPrenda).Insert(c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//Modificar datos de una prenda en la base
func (c *Prenda) Modificar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionPrenda).UpdateId(c.ID, c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//ConsultarPrendas regresa un catálogo de prendas
func ConsultarPrendas() (prendas []Prenda) {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionPrenda).Find(bson.M{}).All(&prendas)
	if err != nil {
		log.RegistrarError(err)
	}
	return prendas
}

//BuscarPorID busca en la BD un color que coincida con el ID dado
func (c *Prenda) BuscarPorID() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionPrenda).FindId(c.ID).One(c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}
