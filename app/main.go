package main

import (
	"github.com/gherbust/lab/internal/directory/applications"
	"github.com/gherbust/lab/internal/directory/infrastructure"

	//mongoInfrastructure "github.com/gherbust/lab/platform/mongo/infrastructure"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	/*conectionString := os.Getenv("connectionString")
	sqlDB, err := mysqlInfrastructure.OpenMysqlDB(conectionString)
	if err != nil {
		panic(err)
	}

	url := "mongodb://localhost:27017"
	mongoClient, err := mongoInfrastructure.OpenMongoDB(url)
	if err != nil {
		fmt.Println(err.Error())
		panic("Error al abrir mongo")
	}
	defer mongoClient.Disconnect(context.Background())
	*/
	directory := applications.NewDirectory()
	handler := infrastructure.NewDirectoryHandler(directory)
	r := gin.Default()
	r.LoadHTMLGlob("../templates/*.html")

	r.POST("/directory", handler.Create)
	r.GET("/directory/:name", handler.GetByName)
	r.GET("/directory", handler.GetAll)

	r.GET("/", handler.Contact)
	views := r.Group("/contact")
	views.GET("/", handler.ContactDetail)
	views.GET("/create", handler.ViewContactCreate)
	views.POST("/create", handler.ContactCreate)

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
