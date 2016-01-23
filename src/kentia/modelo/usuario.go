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
}
u
