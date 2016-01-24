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
	Genotipo []modelo.FormaColor
	Aptitud  float64
}

func crearIndividuo(cp modelo.ColoresPrendas) (ind Individuo) {
	ind.Genotipo = make([]modelo.FormaColor, 4)
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
	h1.Genotipo, h2.Genotipo = make([]modelo.FormaColor, 4), make([]modelo.FormaColor, 4)
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

func relacionarTonos(t1, t2 int) int {
	if t1 == t2 {
		return monocromatico
	}
	if (t1+1)%12 == t2 || (t1-1)%12 == t2 {
		return analogo
	}
	if (t1+6)%12 == t2 {
		return complementario
	}
	if t2 > 11 {
		return comodin
	}
	return sinRelacion
}

func relacionarBrillos(b1, b2 int) bool {
	if b1 == b2 || b1 == (b2+1)%5 || b1 == (b2-1)%5 {
		return true
	}
	return false
}

func calcularAptitud(contador map[int]int, contadorBrillo int) (apt float64) {
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
	apt *= float64(contadorBrillo)/6 + 1
	return apt
}

func (ind *Individuo) evaluar() {
	for i := range ind.Genotipo {
		contador := make(map[int]int)
		contadorBrillo := 0
		if ind.Genotipo[i].Tono < 12 {
			for j := 1; j < 4; j++ {
				sig := (i + j) % 4
				contador[relacionarTonos(ind.Genotipo[i].Tono, ind.Genotipo[sig].Tono)]++
				if relacionarBrillos(ind.Genotipo[i].Brillo, ind.Genotipo[sig].Brillo) {
					contadorBrillo++
				}
			}
		}
		nuevaAptitud := calcularAptitud(contador, contadorBrillo)
		if nuevaAptitud > ind.Aptitud {
			ind.Aptitud = nuevaAptitud
		}
	}
}
