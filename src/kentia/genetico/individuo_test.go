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
func TestCrearIndividuo(t *testing.T) {
	cp := modelo.ColoresPrendas{}
	cp.Calzado = []int{2, 5, 4, 2, 3}
	cp.Chamarra = []int{2, 3, 5, 3, 1, 2}
	cp.Pantalon = []int{2, 3, 1, 3, 3}
	cp.Playera = []int{2, 3, 1, 3, 1, 11}

	ind1 := crearIndividuo(cp)
	/*ind2 := crearIndividuo(cp)
	fmt.Println(ind1, ind2)
	fmt.Println(ind1.cruza(ind2))*/
	ind1.mutar(cp)
	fmt.Println(ind1)
}
