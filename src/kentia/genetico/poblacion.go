package genetico

import (
	"fmt"
	"kentia/modelo"
	"math/rand"
	"sort"
)

const (
	individuos   = 100
	generaciones = 10
	pm           = .1
)

type poblacion []Individuo

func (p poblacion) Len() int {
	return len(p)
}

func (p poblacion) Less(i, j int) bool {
	return p[i].Aptitud > p[j].Aptitud
}

func (p poblacion) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func ordenar(p *poblacion) {
	sort.Sort(p)
}

func crearPoblacion(cp modelo.ColoresPrendas) (pob poblacion) {
	for i := 0; i < individuos; i++ {
		pob = append(pob, crearIndividuo(cp))
	}
	return pob
}

func (p poblacion) mutarEvaluar(cp modelo.ColoresPrendas) poblacion {
	for i := range p {
		prob := rand.Float64()
		if prob <= pm {
			p[i].mutar(cp)
		}
		p[i].evaluar()
	}
	return p
}

func (p poblacion) seleccion() Individuo {
	totalAptitutdes := 0.0
	for _, ind := range p {
		totalAptitutdes += ind.Aptitud
	}
	esperanzas := make([]float64, len(p))
	for i := range esperanzas {
		esperanzas[i] = p[i].Aptitud / totalAptitutdes
	}
	i := 0
	pa := esperanzas[i]
	ultimoIndice := len(esperanzas) - 1
	random := rand.Float64()

	for random > pa {
		i++
		if i == ultimoIndice {
			break
		}
		pa += esperanzas[i]
	}
	return p[i]
}

func (p poblacion) elegirMejores() poblacion {
	ordenar(&p)
	return p[:individuos]
}

func (p poblacion) crearHijos() (hijos poblacion) {
	for i := 0; i < len(p)/2; i++ {
		padre1, padre2 := p.seleccion(), p.seleccion()
		h1, h2 := padre1.cruza(padre2)
		hijos = append(hijos, h1, h2)
	}
	return hijos
}

//Genetico genera el algoritmo génetico para combinar colores.
func Genetico(cp modelo.ColoresPrendas) []Individuo {
	pob := crearPoblacion(cp)
	ordenar(&pob)
	for i := 0; i < generaciones; i++ {
		hijos := pob.crearHijos()
		hijos = hijos.mutarEvaluar(cp)
		pob = append(pob, hijos...)
		pob = pob.elegirMejores()
		fmt.Println("\nMejor generacion ", (i + 1))
		fmt.Println("Genotipo: ", pob[0].Genotipo, " Aptitud total: ", pob[0].Aptitud)
	}
	return pob[:3]
}
