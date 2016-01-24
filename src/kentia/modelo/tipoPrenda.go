package modelo

import (
	"kentia/log"

	"gopkg.in/mgo.v2/bson"
)

//TipoPrenda es la estructura que define los climas para los que se usara la prenda
type TipoPrenda struct {
	ID     bson.ObjectId `bson:"_id" form:"tiposPrenda" binding:"required"`
	Nombre string
}

const coleccionTipoPrenda = "tipo_prenda"

//Registrar se encarga de registrar el clima en la BD
func (tp *TipoPrenda) Registrar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionTipoPrenda).Insert(tp)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//ConsultarPorID para consultar
func (tp *TipoPrenda) ConsultarPorID() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionTipoPrenda).FindId(tp.ID).One(tp)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//ConsultarTiposPrenda consulta todas los tipos de prenda
func ConsultarTiposPrenda() (tiposPrenda []TipoPrenda) {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionTipoPrenda).Find(bson.M{}).All(&tiposPrenda)
	if err != nil {
		log.RegistrarError(err)
	}
	return tiposPrenda
}

//Modificar se encarga de modificar la combinacion en la BD
func (tp *TipoPrenda) Modificar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionTipoPrenda).UpdateId(tp.ID, tp)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}
