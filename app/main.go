package main

import (
	"github.com/gherbust/lab/internal/shared"
)

func main() {
	// shared.Multiplicacion(889)

	// lista := []int{3, 4, 56, 45, 23, 89, 97, 567, 455, 334, 3345, 2, 4, 6}
	// shared.ObtenerLimites(lista)
	// contraseña := shared.CrearContraseña(8)
	// fmt.Println(contraseña)

	// if shared.EsPalindromo(" amor a Roma ") {
	// 	fmt.Println("Es palindromo")
	// } else {
	// 	fmt.Println("No es palindromo")
	// }

	// s := shared.CerdoLatinoFrace(" the big grean apple")
	// fmt.Println(s)

	// shared.Bucle()
	nombres := []string{"Jose", "Ricardo", "Pablo", "Andres", "Juan", "Daniel", "Bernardo", "Claudia"}

	shared.OrdenarNombres(&nombres)

	// for _, v := range nombres {
	// 	fmt.Println(v)
	// }
}
