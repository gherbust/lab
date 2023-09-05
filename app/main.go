package main

import (
	"github.com/gherbust/lab/internal/directory/applications"
	"github.com/gherbust/lab/internal/directory/infrastructure"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
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

	r.Run()
}
