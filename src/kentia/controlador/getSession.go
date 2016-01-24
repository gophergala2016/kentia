package controlador

import "gopkg.in/mgo.v2/bson"

/*GetSessionBsonObjectID obtiene el valor bson.ObjectId de la sesion actual*/
func GetSessionBsonObjectID(session interface{}) bson.ObjectId {
	var mySession bson.ObjectId
	switch session.(type) {
	default:
		mySession = bson.ObjectId(0)
	case bson.ObjectId:
		mySession = session.(bson.ObjectId)
	}
	return mySession
}
