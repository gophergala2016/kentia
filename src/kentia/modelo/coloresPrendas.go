package modelo

import "math/rand"

//ColoresPrendas indica los colores de las prendas disponibles
type ColoresPrendas struct {
	Calzado  []FormaColor
	Pantalon []FormaColor
	Playera  []FormaColor
	Chamarra []FormaColor
}

//FormaColor informacion básica de un color para una prenda
type FormaColor struct {
	Tono   int
	Brillo int
}

//GetColores regresa los colores de una prenda en específico
func (p ColoresPrendas) GetColores(n int) []FormaColor {
	switch n {
	case 0:
		return p.Calzado
	case 1:
		return p.Pantalon
	case 2:
		return p.Playera
	case 3:
		return p.Chamarra
	default:
		return nil
	}
}

//GetRandom regresa un elemento random para el elmento n
func (p ColoresPrendas) GetRandom(n int) FormaColor {
	disponibles := p.GetColores(n)
	return disponibles[rand.Intn(len(disponibles))]
}

//ConsultarColoresPrendas regresa los colores de prenda que tiene el usuario
func ConsultarColoresPrendas(usuarioID string) (cp ColoresPrendas) {
	return cp
}
