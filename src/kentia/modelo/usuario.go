package modelo

import (
	"kentia/log"

	"gopkg.in/mgo.v2/bson"
)

//Usuario define los valores que identifican a un usario del sistema
type Usuario struct {
	ID         bson.ObjectId `bson:"_id"`
	Nombre     string        `form:"nombre"`
	Correo     string        `form:"correo"`
	Contraseña string        `form:"pass"`
	Genero     string        `form:"genero"`
	Prendas    []Prenda
}

const coleccionUsuario = "usuario"

//Registrar registra un usuario en la DB
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

//Modificar modifica un usuario en la DB
func (u *Usuario) Modificar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionUsuario).UpdateId(u.ID, u)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//ConsultarUsuarios regresa el catálogo de usuarios
func ConsultarUsuarios() (usuarios []Usuario) {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionUsuario).Find(bson.M{}).All(&usuarios)
	if err != nil {
		log.RegistrarError(err)
	}

	return usuarios
}

//BuscarPorID busca un usuario en la DB por ID
func (u *Usuario) BuscarPorID() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.C(coleccionUsuario).Find(u.ID).One(u)
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}
