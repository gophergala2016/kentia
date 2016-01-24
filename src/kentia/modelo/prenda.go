package modelo

import (
	"kentia/log"

	"gopkg.in/mgo.v2/bson"
)

//Prenda define los datos importantes para una prenda.
type Prenda struct {
	ID         bson.ObjectId `bson:"_id"`
	Brillo     int           `form:"brillo"`
	Foto       string
	Color      Color
	Clima      Clima
	TipoPrenda TipoPrenda
	Ocasion    Ocasion
}

const coleccionPrenda = "prenda"

//Registrar inserta la prenda en BD.
func (p *Prenda) Registrar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionPrenda).Insert(p)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//Modificar datos de una prenda en la base
func (p *Prenda) Modificar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionPrenda).UpdateId(p.ID, p)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//ConsultarPrendas regresa un catálogo de prendas.
func ConsultarPrendas() (prendas []Prenda) {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionPrenda).Find(bson.M{}).All(&prendas)
	if err != nil {
		log.RegistrarError(err)
	}
	return prendas
}

//BuscarPorID busca en la BD un color que coincida con el ID dado.
func (p *Prenda) BuscarPorID() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionPrenda).FindId(p.ID).One(p)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//BuscarPorBrilloTono busca en la BD un color que coincida con el tono y el brillo.
func (p *Prenda) BuscarPorBrilloTono(prendas []Prenda) bool {
	for _, prenda := range prendas {
		if p.TipoPrenda.Nombre == prenda.TipoPrenda.Nombre && p.Color.Tono == prenda.Color.Tono && p.Brillo == prenda.Brillo {
			*p = prenda
			return true
		}
	}
	return false
}
