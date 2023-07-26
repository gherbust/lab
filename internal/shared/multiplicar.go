package shared

import "fmt"

func Multiplicacion(valor int) {
	for i := 0; i <= 10; i++ {
		r := valor * i
		fmt.Printf("%v X %v = %v \n", valor, i, r)
	}
}
