package genetico

import (
	"fmt"
	"kentia/modelo"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestCrearPoblacion(t *testing.T) {
	prendas := 20
	cp := modelo.ColoresPrendas{}
	cp.Calzado = make([]modelo.FormaColor, prendas)
	cp.Chamarra = make([]modelo.FormaColor, prendas)
	cp.Pantalon = make([]modelo.FormaColor, prendas)
	cp.Playera = make([]modelo.FormaColor, prendas)
	for i := 0; i < prendas; i++ {
		cp.Calzado[i].Tono = rand.Intn(prendas)
		cp.Chamarra[i].Tono = rand.Intn(prendas)
		cp.Pantalon[i].Tono = rand.Intn(prendas)
		cp.Playera[i].Tono = rand.Intn(prendas)
	}
	mejores := Genetico(cp)
	fmt.Println("\nLo mejor de lo mejor de lo mejor de lo mejor de lo mejor de lo mejor de lo mejor de lo mejor:\n", mejores[0], "\n", mejores[1], "\n", mejores[2])
}
