package main

import (
	"fmt"

	"github.com/gherbust-meli/lab/internal/shared"
)

func main() {

	nombres := []string{"Jose", "Ricardo", "Pablo", "Mia", "Lunita", "Gorda", "Bbeshito", "HijoPanzon"}
	shared.OrdenarNombres(&nombres)

	for _, v := range nombres {
		fmt.Println(v)
	}
	/*

		shared.NumerosB(10)
	*/
	/*
		s := shared.CerdoLatinoFrase("Hola mundo jeje patata")
		fmt.Println(s)
	*/
	/*
		s := shared.CerdoLatinoPalabra("Hello")
		fmt.Println(s)
	*/
	/*
		if shared.EsPalindromo("AmofRoma") {
			fmt.Println("Es palindromo")
		} else {
			fmt.Println("No es palindromo")
		}
	*/
	/*
		contraseña := shared.CrearContraseña(12)
		fmt.Println(contraseña)

		v := "sakdfjsdgnlkjkmkmmvf"

		for i, v := range v {
			fmt.Println(i)
			fmt.Println(v)
		}
	*/
	/*
		//shared.Multiplicacion(898)
		//lista := []int{3, 4, 56, 45, 23, 89, 97, 527}
		//shared.ObtenerLimites(lista)
		/*directory := applications.NewDirectory()
		handler := infrastructure.NewDirectoryHandler(*directory)
		r := gin.Default()
		r.POST("/directory", handler.Create)
		r.GET("/directory/:name", handler.GetByName)
		r.GET("/directory", handler.GetAll)

		r.POST("/stringConverter", stringfuntionsinfraestructure.StringConverter)

		r.Run()
	*/

	/*	c := domain.Contact{
			Name:        "Dul",
			PhoneNumber: "5587654321",
			EMail:       "dul@correo.com",
		}

		directory := applications.NewDirectory()

		directory.SaveContact(c)

		contact := directory.GetContact("Dul")

		fmt.Println(contact.EMail)
	*/

	/*	var edad int
		edad = 21
		año := 1983
		edad, err := calculaEdad(año)

		if err != nil {
			fmt.Println(err.Error())NumerosB(numeros int)OrdenarNombres(nombres *[]string)
			panic(err.Error())
		}

		msg := fmt.Sprintf("%v edad: %v nacio en el ano %v", "gerardo", edad, año)
		fmt.Println(msg)
	*/

}

/*
func calculaEdad(añoNacimiento int) (int, error) {
	año := time.Now().Year()
	edad := año - añoNacimiento
	return edad, nil
}*/
