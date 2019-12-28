package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    true,
		"websocket": "connect",
	})
}
