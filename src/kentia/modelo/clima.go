package modelo

import "gopkg.in/mgo.v2/bson"

//Clima tipo de clima para el que se usa esta prenda
type Clima struct {
	ID     bson.ObjectId `bson:"_id" form:"id"`
	Nombre string
}

const coleccionClima = "clima"

//Registar funcion para cargar clima
func (c *Clima) Registar() {
	var conn conector
	conn.IniciarSesion()
	defer conn.CerrarSesion()
err:
	dao.db.C(coleccionClima).Insert(c)
	if err != nil {
		lod.RegistarError(err)
		return false
	}
	return true
}
