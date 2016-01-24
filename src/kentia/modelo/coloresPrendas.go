package modelo

import "math/rand"

//ColoresPrendas indica los colores de las prendas disponibles
type ColoresPrendas struct {
	Calzado  []int
	Pantalon []int
	Playera  []int
	Chamarra []int
}

//GetColores regresa los colores de una prenda en específico
func (p ColoresPrendas) GetColores(n int) []int {
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

func (p ColoresPrendas) GetRandom(n int) int {
	disponibles := p.GetColores(n)
	return disponibles[rand.Intn(len(disponibles))]
}
