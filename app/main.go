package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gherbust/lab/internal/directory/applications"
	"github.com/gherbust/lab/internal/directory/infrastructure"
	"github.com/gherbust/lab/internal/platform/mysql/domain"
	stringfuntionsinfraestructure "github.com/gherbust/lab/internal/stringfuntions/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := sql.Open("mysql", "root:bebes2023*@tcp(localhost:3306)/directorio") //aqui hay datos que vemos en la configuracion de la base de datos "edit conection"

	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	query := "INSERT INTO `directorio`.`contacto`(`name`,`phone_number`,`e_mail`,`enabled`,`last_update`) values(?,?,?,?,?)"

	result, err := db.Exec(query, "Dulce5", "8383455387522", "dulce5@correo.com", 1, "2023-07-31 17:14:00")
	if err != nil {
		fmt.Println(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)

	query = "SELECT idcontacto,name,phone_number,e_mail,enabled FROM directorio.contacto"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	contactos := []domain.Contact{} //se ponen llavesitas para instanciar el objeto "{}"
	for rows.Next() {
		contacto := new(domain.Contact)
		rows.Scan(&contacto.Id, &contacto.Name, &contacto.PhoneNumber, &contacto.EMail, &contacto.Enabled)
		contactos = append(contactos, *contacto)
	}

	fmt.Println(len(contactos))

	/*
		nombres := []string{"Jose", "Ricardo", "Pablo", "Mia", "Lunita", "Gorda", "Bebeshito", "HijoPanzon"}
		shared.OrdenarNombres(&nombres)

		for _, v := range nombres {
			fmt.Println(v)
		}
	*/
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
	*/
	directory := applications.NewDirectoryMYSQL()
	handler := infrastructure.NewDirectoryHandler(directory)
	r := gin.Default()
	r.POST("/directory", handler.Create)
	r.GET("/directory/:name", handler.GetByName)
	r.GET("/directory", handler.GetAll)

	r.POST("/stringConverter", stringfuntionsinfraestructure.StringConverter)

	r.Run()

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
