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
