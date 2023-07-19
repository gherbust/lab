package infrastructure

import (
	"net/http"

	"github.com/gherbust-meli/lab/internal/applications"
	"github.com/gherbust-meli/lab/internal/domain"
	"github.com/gin-gonic/gin"
)

type DirectoryHandler struct {
	Directory applications.Directory
}

func NewDirectoryHandler(directory applications.Directory) *DirectoryHandler {
	return &DirectoryHandler{
		Directory: directory,
	}
}

func (d *DirectoryHandler) GetByName(c *gin.Context) {
	name := c.Param("name")
	contact := d.Directory.GetContact(name)
	c.JSON(http.StatusAccepted, contact)

}

func (d *DirectoryHandler) Create(c *gin.Context) {
	request := domain.Contact{}
	c.Bind(&request)
	d.Directory.SaveContact(request)
	c.JSON(http.StatusAccepted, nil)

}
