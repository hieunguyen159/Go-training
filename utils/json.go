package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ERROR(c *gin.Context, statusCode int, err error) {
	if err != nil {
		c.JSON(statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusBadRequest, nil)
}
