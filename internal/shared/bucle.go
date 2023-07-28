package shared

import "fmt"

///crea un programa en bucle que imprima los numeros del 1 al 10///

func NumerosB(numeros int) {
	for i := 0; i < numeros; i++ {
		n := i + 1
		fmt.Printf("%v \n", n)
	}
}

///lista de nombres que ordene alfabeticamente
