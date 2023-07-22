package main

import (
	"github.com/gherbust-meli/lab/internal/applications"
	"github.com/gherbust-meli/lab/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	directory := applications.NewDirectory()
	handler := infrastructure.NewDirectoryHandler(*directory)
	r := gin.Default()
	r.POST("/directory", handler.Create)
	r.GET("/directory/:name", handler.GetByName)
	r.GET("/directory", handler.GetALL)

	r.Run()

	/*c := domain.Contact{
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
			fmt.Println(err.Error())
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
