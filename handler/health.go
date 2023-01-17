package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	res := []string{"aman", "gupta"}
	c.JSON(http.StatusOK, res)
}
