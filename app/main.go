package main

import (
	"github.com/gherbust/lab/internal/directory/applications"
	"github.com/gherbust/lab/internal/directory/infrastructure"
	"github.com/gherbust/lab/internal/shared"
	stringsfunctionsinfrastructure "github.com/gherbust/lab/internal/stringsfunctions/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	directory := applications.NewDirectoryMYSQL()
	handler := infrastructure.NewDirectoryHandler(directory)
	r := gin.Default()
	r.POST("/directory", handler.Create)
	r.GET("/directory/:name", handler.GetByName)
	r.GET("/directory", handler.GetAll)

	r.POST("/stringConverter", stringsfunctionsinfrastructure.StringConverter)

	r.Run()

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
