package controlador

import (
	"kentia/genetico"
	"kentia/modelo"
)

//GenerarMejorCombinacion s
func GenerarMejorCombinacion(usuarioID string) (prendas [][]modelo.Prenda) {
	mejores := genetico.Genetico(modelo.ConsultarColoresPrendas(usuarioID))
	for _, mejor := range mejores {
		var combinacion []modelo.Prenda
		for _, color := range mejor.Genotipo {
			prenda := modelo.Prenda{}
			prenda.Brillo = color.Brillo
			prenda.Color.Tono = color.Tono
			prenda.ConsularPorTonoBrillo()
			combinacion = append(combinacion, prenda)
		}
		prendas = append(prendas, combinacion)
	}
	return prendas
}
