package infrastructure

import (
	"net/http"

	"github.com/gherbust/lab/internal/directory/applications"
	"github.com/gherbust/lab/internal/directory/domain"
	"github.com/gin-gonic/gin"
)

type DirectoryHandler struct {
	Directory applications.DirectoryRepository
}

func NewDirectoryHandler(directory applications.DirectoryRepository) *DirectoryHandler {
	return &DirectoryHandler{
		Directory: directory,
	}
}

func (d *DirectoryHandler) GetByName(c *gin.Context) {
	name := c.Param("name")
	contact := d.Directory.GetContact(name)
	c.JSON(http.StatusAccepted, contact)

}

func (d *DirectoryHandler) GetAll(c *gin.Context) {
	contacts := d.Directory.GetAll()
	c.JSON(http.StatusOK, contacts)

}

func (d *DirectoryHandler) Create(c *gin.Context) {
	request := domain.Contact{}
	c.Bind(&request)
	d.Directory.SaveContact(request)
	c.JSON(http.StatusAccepted, nil)

}

func (d *DirectoryHandler) Contact(c *gin.Context) {
	contacts := d.Directory.GetAllEnabled()
	data := gin.H{
		"title":    "Directorio",
		"name":     "Gerardo HB",
		"contacts": contacts,
	}
	c.HTML(http.StatusOK, "index.html", data)

}
