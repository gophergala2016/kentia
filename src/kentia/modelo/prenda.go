package modelo

import "gopkg.in/mgo.v2/bson"

//Clima tipo de clima para el que se usa esta prenda
type Prenda struct {
	ID         bson.ObjectId `bson:"_id" form:"id"`
	Tono       int
	Foto       string
	color      Color
	clima      Clima
	tipoprenda TipoPrenda
	ocasion    Ocasion
	usuario    Usuario
}

const coleccionPrenda = "prenda"

//Registar funcion para cargar clima
func (c *Prenda) Registar() {
	var conn conector
	conn.IniciarSesion()
	defer conn.CerrarSesion()
err:
	dao.db.C(coleccionPrenda).Insert(c)
	if err != nil {
		lod.RegistarError(err)
		return false
	}
	return true
}
