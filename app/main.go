package main

import (
	"os"

	"github.com/gherbust/lab/internal/directory/applications"
	"github.com/gherbust/lab/internal/directory/infrastructure"
	mysqlInfrastructure "github.com/gherbust/lab/platform/mysql/infrastructure"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conectionString := os.Getenv("connectionString")
	sqlDB, err := mysqlInfrastructure.OpenMysqlDB(conectionString)
	if err != nil {
		panic(err)
	}

	directory := applications.NewDirectoryMYSQL(sqlDB)
	handler := infrastructure.NewDirectoryHandler(directory)
	r := gin.New()
	r.LoadHTMLGlob("../templates/*.html")
	apiRouts := r.Group("/api")
	apiRouts.POST("/directory", handler.Create)
	apiRouts.GET("/directory/:name", handler.GetByName)
	apiRouts.GET("/directory", handler.GetAll)

	r.GET("/", handler.TemplateGetAll)
	r.GET("/contact", handler.TemplateGetByName)
	r.GET("/contact/create", handler.TempateContactCreate)
	r.POST("/contact/create", handler.TemplateCreate)

	//r.POST("/stringConverter", stringsfunctionsinfrastructure.StringConverter)

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
	// nombres := []string{"Jose", "Ricardo", "Pablo", "Andres", "Juan", "Daniel", "Bernardo", "Claudia"}

	// shared.OrdenarNombres(&nombres)

	// for _, v := range nombres {
	// 	fmt.Println(v)
	// }
}
