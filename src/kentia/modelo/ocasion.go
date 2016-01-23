package modelo

import "gopkg.in/mgo.v2/bson"

//tipo de clima para el que se usa esta prenda
type Ocasion struct{
  ID         bson.ObjectId `bson:"_id" form:"id"`
  Nombre string
}

const coleccionOcasion = "ocacion"

func (c *Ocasion)Registar(){
  car conn conector
  conn.IniciarSesion()
  defer conn.CerrarSesion()
  err :dao.db.C(coleccionOcacion).Insert(c)
  if err != nil{
    lod.RegistarError(err)
    return false
  }
  return true
}
