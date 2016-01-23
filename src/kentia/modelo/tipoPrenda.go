package modelo

import (
	"kentia/log"

	"gopkg.in/mgo.v2/bson"
)

//tipoPrenda es la estructura que define los climas para los que se usara la prenda
type tipoPrenda struct {
	ID     bson.ObjectId `bson:"_id" form:"id"`
	Nombre string
}

const colecciontipoPrenda = "tipo_prenda"

//Registrar se encarga de registrar el clima en la BD
func (tp *tipoPrenda) Registrar() bool {
	var conn conector
	conn.IniciarSesion()
	defer conn.CerrarSesion()
	err := conn.db.C(colecciontipoPrenda).Insert(tp)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}
