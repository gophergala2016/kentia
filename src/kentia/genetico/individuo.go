package genetico

import (
	"kentia/modelo"
	"math"
	"math/rand"
)

const (
	phi            = 1.6180339887
	sinRelacion    = 0
	analogo        = 1
	complementario = 2
	monocromatico  = 3
	comodin        = 4
)

//Individuo es el
type Individuo struct {
	Genotipo []int
	Aptitud  float64
}

func crearIndividuo(cp modelo.ColoresPrendas) (ind Individuo) {
	ind.Genotipo = make([]int, 4)
	for i := 0; i < 4; i++ {
		ind.Genotipo[i] = cp.GetRandom(i)
	}
	ind.evaluar()
	return ind
}

func (ind *Individuo) mutar(cp modelo.ColoresPrendas) {
	p := rand.Intn(4)
	ind.Genotipo[p] = cp.GetRandom(p)
}

func (ind *Individuo) cruza(pareja Individuo) (h1, h2 Individuo) {
	h1.Genotipo, h2.Genotipo = make([]int, 4), make([]int, 4)
	copy(h1.Genotipo, ind.Genotipo)
	copy(h2.Genotipo, pareja.Genotipo)
	numCambios := rand.Intn(3) + 1
	posCambios := make([]int, numCambios)
	for i := 0; i < numCambios; i++ {
		posCambios[i] = rand.Intn(4)
		for j := i - 1; j >= 0; j-- {
			if posCambios[i] == posCambios[j] {
				i--
				break
			}
		}
	}

	for pos := range posCambios {
		h2.Genotipo[pos], h1.Genotipo[pos] = h1.Genotipo[pos], h2.Genotipo[pos]
	}
	return h1, h2
}

func relacionarColores(c1, c2 int) int {
	if c1 == c2 {
		return monocromatico
	}
	if (c1+1)%12 == c2 || (c1-1)%12 == c2 {
		return analogo
	}
	if (c1+6)%12 == c2 {
		return complementario
	}
	if c2 > 11 {
		return comodin
	}
	return sinRelacion
}

func calcularAptitud(contador map[int]int) (apt float64) {
	mayor := 1
	for i := 2; i <= 3; i++ {
		if contador[i] > mayor {
			mayor = i
		}
	}
	contador[mayor] += contador[comodin]
	apt = math.Pow(phi, float64(2*contador[analogo]))
	apt += math.Pow(phi, float64(2*contador[complementario]))
	apt += math.Pow(phi, float64(2*contador[monocromatico]))
	return apt
}

func (ind *Individuo) evaluar() {
	for i := range ind.Genotipo {
		contador := make(map[int]int)
		if ind.Genotipo[i] < 12 {
			for j := 1; j < 4; j++ {
				sig := (i + j) % 4
				contador[relacionarColores(ind.Genotipo[i], ind.Genotipo[sig])]++
			}
		}
		nuevaAptitud := calcularAptitud(contador)
		if nuevaAptitud > ind.Aptitud {
			ind.Aptitud = nuevaAptitud
		}
	}
}
