package controlador

import (
	"kentia/genetico"
	"kentia/modelo"

	"gopkg.in/mgo.v2/bson"
)

//GenerarMejorCombinacion se encarga de buscar cada una de las prendas por color y birllo para generar una combinacion.
func GenerarMejorCombinacion(usuarioID string) (prendas [][]modelo.Prenda) {
	u := modelo.Usuario{ID: bson.ObjectIdHex(usuarioID)}
	u.BuscarPorID()
	mejores := genetico.Genetico(u.ConsultarColoresPrendas())
	for _, mejor := range mejores {
		var combinacion []modelo.Prenda
		for i, color := range mejor.Genotipo {
			prenda := modelo.Prenda{}
			prenda.Brillo = color.Brillo
			prenda.Color.Tono = color.Tono
			switch i {
			case 0:
				prenda.TipoPrenda.Nombre = "Calzado"
			case 1:
				prenda.TipoPrenda.Nombre = "Pantalon/Falda"
			case 2:
				prenda.TipoPrenda.Nombre = "Playera"
			case 3:
				prenda.TipoPrenda.Nombre = "Chamarra"
			}
			prenda.BuscarPorBrilloTono(u.Prendas)
			combinacion = append(combinacion, prenda)
		}
		prendas = append(prendas, combinacion)
	}
	return prendas
}
