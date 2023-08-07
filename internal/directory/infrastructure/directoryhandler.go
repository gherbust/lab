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

func (d *DirectoryHandler) ContactDetail(c *gin.Context) {
	name := c.Query("name")
	contact := d.Directory.GetContact(name)
	data := gin.H{
		"title":   "Contacto",
		"contact": contact,
	}

	c.HTML(http.StatusOK, "contact.html", data)
}

func (d *DirectoryHandler) ViewContactCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "new_contact.html", nil)
}

func (d *DirectoryHandler) ContactCreate(c *gin.Context) {
	contact := domain.Contact{}
	contact.Name = c.PostForm("name")
	contact.PhoneNumber = c.PostForm("phone_number")
	contact.EMail = c.PostForm("e_mail")

	d.Directory.SaveContact(contact)
	c.HTML(http.StatusOK, "new_contact.html", nil)
}
