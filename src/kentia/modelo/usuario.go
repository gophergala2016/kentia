package modelo

import (
	"gopkg.in/mgo.v2/bson"
	"kentia/log"
)

type Usuario struct {
	ID         bson.ObjectId `bson:"_id"`
	Nombre     string
	Correo     string
	Contraseña string
	Genero     string
	Prendas    []Prenda
}

const coleccionUsuario = "usuario"

//Registrar. Registra un usuario en la DB
func (u *Usuario) Registrar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionUsuario).Insert(u)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//Modificar. Modifica un usuario en la DB
func (u *Usuario) Modificar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionUsuario).UpdateId(c.ID, c)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//Consultar. Regresa el catálogo de usuarios
func ConsultarUsuarios() (usuarios []Usuario) {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionUsuario).Find(bson.M{}).All(&usuarios)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//BuscarPorID. Busca un usuario en la DB por ID
func (u *Usario) BuscarPorID() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionUsuario).Find(u.ID).One(u)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}
