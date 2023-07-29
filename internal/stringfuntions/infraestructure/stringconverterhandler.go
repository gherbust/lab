package infraestructure

import (
	"net/http"

	"github.com/gherbust/lab/internal/stringfuntions/applications"
	"github.com/gin-gonic/gin"
)

func StringConverter(c *gin.Context) {
	request := Request{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, "payloud con error")
		return
	}

	newText := applications.StringConverter(request.Action, request.Text)

	response := Response{
		OriginalText: request.Text,
		NewText:      newText,
	}

	c.JSON(http.StatusOK, response)

}
