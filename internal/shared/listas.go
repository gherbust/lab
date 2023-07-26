package shared

import (
	"fmt"
	"sort"
)

func ObtenerLimites(lista []int) {
	sort.Sort(sort.IntSlice(lista))
	menor := lista[0]
	ultimoIndice := len(lista) - 1
	mayor := lista[ultimoIndice]

	fmt.Printf("Menor: %v Mayor: %v", menor, mayor)
}
