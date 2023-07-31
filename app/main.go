package main

import (
	"database/sql"
	"fmt"

	"github.com/gherbust/lab/platform/mysql/domain"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:Subreg03@tcp(localhost:3306)/directorio")
	//aqui hay datos que vemos en la configuracion de la base de datos "edit conection"

	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	query := "INSERT INTO `directorio`.`contacto`(`name`,`phone_number`,`e_mail`,`enabled`,`last_update`) values(?,?,?,?,?)"

	result, err := db.Exec(query, "Dulce1", "5587654314", "dul1@correo.com", 1, "2023-07-30 00:22:00")
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
	contactos := []domain.Contact{}
	for rows.Next() {
		contacto := new(domain.Contact)
		err = rows.Scan(&contacto.Id, &contacto.Name, &contacto.PhoneNumber, &contacto.EMail, &contacto.Enabled)
		if err != nil {
			fmt.Println(err.Error())
		}
		contactos = append(contactos, *contacto)
	}

	fmt.Println(len(contactos))

	/*
		directory := applications.NewDirectoryMYSQL()
		handler := infrastructure.NewDirectoryHandler(directory)
		r := gin.Default()
		r.POST("/directory", handler.Create)
		r.GET("/directory/:name", handler.GetByName)
		r.GET("/directory", handler.GetAll)

		r.POST("/stringConverter", stringsfunctionsinfrastructure.StringConverter)

		r.Run()
	*/

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
	// nombres := []string{"Jose", "Ricardo", "Pablo", "Andres", "Juan", "Daniel", "Bernardo", "Claudia"}

	// shared.OrdenarNombres(&nombres)

	// for _, v := range nombres {
	// 	fmt.Println(v)
	// }
}
