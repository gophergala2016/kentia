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
	prendas := 10
	cp := modelo.ColoresPrendas{}
	cp.Calzado = make([]int, 3)
	cp.Chamarra = make([]int, prendas)
	cp.Pantalon = make([]int, prendas)
	cp.Playera = make([]int, prendas)
	for i := 0; i < prendas; i++ {
		cp.Calzado[i] = rand.Intn(12)
		cp.Chamarra[i] = rand.Intn(12)
		cp.Pantalon[i] = rand.Intn(12)
		cp.Playera[i] = rand.Intn(12)
	}
	mejores := Genetico(cp)
	fmt.Println("\nLo mejor de lo mejor de lo mejor de lo mejor de lo mejor de lo mejor de lo mejor de lo mejor:\n", mejores[0], "\n", mejores[1], "\n", mejores[2])
}
